package metrics

const metricName = "pipeline"

// ReportJobCount 上报任务数量，工作队列使用率
func ReportJobCount(jobid uint64, count int64) {

}

// ReportPoolSize 上报工作池大小
func ReportPoolSize(poolSize int64, blockingSize int64) {

}

// ReportSubmitConsume 上报任务提交耗时
func ReportSubmitConsume(consume int64) {

}

// ReportJobConsume 任务消费时间
func ReportJobConsume(jobid uint64, hashkey string, duration int64) {

}

// ReportJobTimeout 任务超时统计
func ReportJobTimeout(jobid uint64, hashkey string) {

}
