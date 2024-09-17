package pipeline

import (
	"fmt"

	"pipeline/dispatcher"
	"pipeline/jobs"
	"pipeline/serial"
)

var (
	// 默认的全局worker queue
	globalWokerQueue = jobs.NewWorkQueue(jobs.GetDefaultConfig())

	// 默认的全局uint64 pipeline
	defaultUint64Pipeline = dispatcher.GetGlobalDispatcher(&serial.Uint64Serializer{})

	// 默认的全局bytes pipeline
	defaultBytesPipeline = dispatcher.GetGlobalDispatcher(&serial.ByteSerializer{})
)

func init() {
	jobs.GlobalWorkerQueueGetter = func() jobs.BaseWorkerQueue {
		return globalWokerQueue
	}
}

// RelaunchWorkerQueue 重置默认队列
func RelaunchDefaultWorkerQueue(cfg *jobs.PipelineConfig) error {
	if cfg == nil ||
		cfg.MaxWorkerQueueCount <= 0 ||
		cfg.MaxJobsPerWorker <= 0 {
		return fmt.Errorf("config is invalid %+v ", cfg)
	}

	if globalWokerQueue != nil {
		globalWokerQueue.Stop()
	}

	globalWokerQueue = jobs.NewWorkQueue(cfg)
	return nil
}

// PostUint64 投递任务
func PostUint64(key uint64, f jobs.Job) error {
	return defaultUint64Pipeline.Post(key, f)
}

// GetJobsBuffLenUint64 获取当前jobs缓冲区长度
func GetJobsBuffLenUint64(key uint64) (int, error) {
	return defaultUint64Pipeline.GetJobsBuffLen(key)
}

// GetQueueIdUint64 获取当前jobID
func GetQueueIdUint64(key uint64) (uint64, error) {
	return defaultUint64Pipeline.GetQueueId(key)
}

// PostBytes 投递任务
func PostBytes(key []byte, f jobs.Job) error {
	return defaultBytesPipeline.Post(key, f)
}

// GetJobsBuffLenBytes 获取当前jobs缓冲区长度
func GetJobsBuffLenBytes(key []byte) (int, error) {
	return defaultBytesPipeline.GetJobsBuffLen(key)
}

// GetQueueIdBytes 获取当前jobID
func GetQueueIdBytes(key []byte) (uint64, error) {
	return defaultBytesPipeline.GetQueueId(key)
}
