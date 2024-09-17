package jobs

// PipelineConfig pipeline 自定义配置
type PipelineConfig struct {
	// 最大工作池大小，默认1000
	MaxWorkerQueueCount int32 `yaml:"max_worker_queue_count"`
	// 每个worker最多处理的任务数，默认10
	// 每个任务队列和worker协程会进行提交绑定，防止任务队列长时间占有worker协程，每次处理一批Job后，将退出绑定，重新提交
	MaxJobsPerWorker int32 `yaml:"max_jobs_per_worker"`
}

// GetDefaultConfig  pipeline 默认数值
func GetDefaultConfig() *PipelineConfig {
	return &PipelineConfig{
		MaxWorkerQueueCount: DefaultMaxWorkerCount,
		MaxJobsPerWorker:    DefaultMaxJobsPerWorker,
	}
}

const (
	DefaultMaxWorkerCount   = 1000 // 最大工作队列
	DefaultMaxJobsPerWorker = 10   // 每个worker最多处理的任务数
)
