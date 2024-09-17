package jobs

import (
	"sync"
	"time"

	"log"
	"pipeline/metrics"
	"pipeline/serial"
	"github.com/panjf2000/ants/v2"
)

// GlobalWorkerQueueGetter 全局工作队列回调
var GlobalWorkerQueueGetter func() BaseWorkerQueue

// job 回调函数
type Job func()

// worker queue
type BaseWorkerQueue interface {
	// 消息派发，消费
	Dispatch(key uint64, f Job)
	// 获取当前jobs缓冲区长度
	JobsBuffLen(key uint64) int
	// 停止
	Stop()
}

type BaseWorker interface {
	// 消费池
	ConsumerPool() *ants.Pool
	// 每批次处理的最多任务数
	MaxJobsPerWorker() int32
}

// JobQueue 任务队列
type JobQueue struct {
	// 任务队列
	jobs *Queue
	// 是否需要提交
	needSubmit bool
	// 全局锁
	sync.Mutex

	// 注入接口
	BaseWorker
}

func (j *JobQueue) equeue(f Job) (isNeedSubmit bool) {
	j.Lock()
	defer j.Unlock()

	j.jobs.Enqueue(f)
	// 首次投递，提交任务
	if j.needSubmit {
		j.needSubmit = false
		return true
	}

	return false
}

func (j *JobQueue) dequeue() Job {
	j.Lock()
	defer j.Unlock()
	f := j.jobs.Dequeue()
	if f != nil {
		return f.(Job)
	}

	return nil
}

// Post 投递任务
func (j *JobQueue) Post(f Job) {
	if j.equeue(f) {
		j.submitTaskBlocking()
	}
}

func (j *JobQueue) needRetrySubmit() (isNeedSubmit bool) {
	j.Lock()
	defer j.Unlock()

	if j.jobs.Size() > 0 {
		// 继续关闭提交开关，返回给调用方立即提交
		j.needSubmit = false
		return true
	} else {
		// 任务队列为空，打开需要提交的开关，等待下次Post来的请求触发提交
		j.needSubmit = true
		return false
	}
}

func (j *JobQueue) doJobs() {
	defer func() {
		// 如果队列中又来了任务，继续提交，这时，post来的job已经跳过了检查提交
		if j.needRetrySubmit() {
			go serial.RecoverGo(j.submitTaskBlocking)
		}
	}()

	for i := int32(0); i < j.MaxJobsPerWorker(); i++ {
		f := j.dequeue()
		if f == nil {
			break
		}
		f()
	}
}

func (j *JobQueue) submitTaskBlocking() {
	metrics.ReportPoolSize(int64(j.ConsumerPool().Running()), int64(j.ConsumerPool().Waiting()))

	now := time.Now()
	if err := j.ConsumerPool().Submit(j.doJobs); err != nil {
		log.Printf("job queue submit error %v pool %d", err, j.ConsumerPool().Running())
	}

	metrics.ReportSubmitConsume(time.Since(now).Milliseconds())
}

func (j *JobQueue) Size() int {
	j.Lock()
	defer j.Unlock()
	return j.jobs.Size()
}

func (j *JobQueue) IsIdle() bool {
	j.Lock()
	defer j.Unlock()
	return j.jobs.Size() == 0 && !j.needSubmit
}

// WorkerQueue	工作队列
type WorkerQueue struct {
	cfg     *PipelineConfig
	stopped bool // 是否停止

	// 消费池
	consumer *ants.Pool

	// 生产池
	provider      map[uint64]*JobQueue // <hashkey, *JobQueue>
	providerMutex sync.Mutex
}

func (w *WorkerQueue) start() (err error) {
	w.consumer, err = ants.NewPool(int(w.cfg.MaxWorkerQueueCount),
		ants.WithPreAlloc(true),
		ants.WithNonblocking(false))
	if err != nil {
		return err
	}

	w.onTimer()
	return
}

// Stop 停止工作队列
func (w *WorkerQueue) Stop() {
	w.consumer.Release()
	w.provider = make(map[uint64]*JobQueue)
	w.stopped = true
}

func (w *WorkerQueue) onTimer() {
	time.AfterFunc(time.Minute, func() {
		w.ClearIdleProvider()

		if !w.stopped {
			w.onTimer()
		}
	})
}

// FetchProvider 获取任务队列
func (w *WorkerQueue) FetchProvider(idx uint64) *JobQueue {
	w.providerMutex.Lock()
	defer w.providerMutex.Unlock()

	queue, ok := w.provider[idx]
	if !ok {
		queue = &JobQueue{
			jobs:       NewQueue(),
			needSubmit: true,
			BaseWorker: w,
		}
		w.provider[idx] = queue
	}

	return queue
}

// ClearIdleProvider 清除空闲队列
func (w *WorkerQueue) ClearIdleProvider() {
	w.providerMutex.Lock()
	defer w.providerMutex.Unlock()

	for idx, queue := range w.provider {
		if queue.IsIdle() {
			delete(w.provider, idx)
		}
	}
}

func (w *WorkerQueue) ConsumerPool() *ants.Pool {
	return w.consumer
}

func (w *WorkerQueue) MaxJobsPerWorker() int32 {
	return w.cfg.MaxJobsPerWorker
}

// Dispatch 任务分发
func (w *WorkerQueue) Dispatch(key uint64, f Job) {
	queue := w.FetchProvider(key)
	queue.Post(f)

	metrics.ReportJobCount(key, int64(queue.Size()))
}

// JobsBuffLen 获取任务队列长度
func (w *WorkerQueue) JobsBuffLen(key uint64) int {
	queue := w.FetchProvider(key)
	return queue.Size()
}

// NewWorkQueue 初始化 worker queue
func NewWorkQueue(cfg *PipelineConfig) BaseWorkerQueue {
	wq := &WorkerQueue{
		cfg:      cfg,
		provider: make(map[uint64]*JobQueue),
	}

	wq.start()
	return wq
}
