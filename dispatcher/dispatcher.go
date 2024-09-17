package dispatcher

import (
	"fmt"

	"pipeline/hash"
	"pipeline/jobs"
)

// hash key 序列化接口
type Serializer[T any] interface {
	Marshal(id T) ([]byte, error)
	Unmarshal(data []byte) (T, error)
}

type PipelineDispatcher[Key any] struct {
	serial      Serializer[Key]
	workerQueue jobs.BaseWorkerQueue
}

// Post 投递消息
func (a *PipelineDispatcher[Key]) Post(id Key, f jobs.Job) error {
	hashvalue, err := a.getHashValue(id)
	if err != nil {
		return err
	}

	worker := a.GetWorkQueue()
	if worker == nil {
		return fmt.Errorf("worker queue is nil")
	}

	worker.Dispatch(hashvalue, f)
	return nil
}

// 使用的 murmur3 算法，计算 hash value
func (a *PipelineDispatcher[Key]) getHashValue(id Key) (uint64, error) {
	idBytes, err := a.serial.Marshal(id)
	if err != nil {
		return 0, err
	}

	return hash.DefaultHashFunc(idBytes, 0)
}

// GetQueueId 根据hashkey获取job id
func (a *PipelineDispatcher[Key]) GetQueueId(id Key) (uint64, error) {
	hashValue, err := a.getHashValue(id)
	if err != nil {
		return 0, err
	}

	return hashValue, nil
}

// GetJobLen 根据hashkey 获取job 缓冲区长度
func (a *PipelineDispatcher[Key]) GetJobsBuffLen(id Key) (int, error) {
	hashValue, err := a.getHashValue(id)
	if err != nil {
		return 0, err
	}

	worker := a.GetWorkQueue()
	if worker == nil {
		return 0, fmt.Errorf("worker queue is nil")
	}

	return worker.JobsBuffLen(hashValue), nil
}

func (a *PipelineDispatcher[Key]) GetWorkQueue() jobs.BaseWorkerQueue {
	if a.workerQueue == nil {
		return jobs.GlobalWorkerQueueGetter()
	}

	return a.workerQueue
}

// GetGlobalDispatcher 获取全局分发器，需要将全局对接回调函数设置给分发器
func GetGlobalDispatcher[Key any](serial Serializer[Key]) *PipelineDispatcher[Key] {
	if serial == nil {
		return nil
	}

	return &PipelineDispatcher[Key]{
		serial: serial,
	}
}

// NewDispatcher 创建分发器，自定义分发器
func NewDispatcher[Key any](serial Serializer[Key], workerQueue jobs.BaseWorkerQueue) *PipelineDispatcher[Key] {
	if serial == nil {
		return nil
	}

	if workerQueue == nil {
		return nil
	}

	return &PipelineDispatcher[Key]{
		workerQueue: workerQueue,
		serial:      serial,
	}
}
