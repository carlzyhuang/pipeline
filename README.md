# actor 处理串行化模块

使用一致性hash key路由的请求，提供一个工作队列池，将请求经过计算hash key的值，投递到工作队列池中，串行消费

## 接入方法

**业务使用方法**
```go
import(
  "pipeline"
)

pipeline.PostBytes(key, func(){

})

pipeline.PostUint64(key, func(){
  
})


```

**注意死锁场景**

业务使用`PostAndWait`投递1号消息并等待，处理1号消息时使用`PostAndWait`产生2号消息并等待

此时，1号消息完成后才会处理2号消息，1号消息又在等待2号消息的返回，产生死锁


## 单元测试

```bash
go test -cover ./...

?       pipeline       [no test files]
?       pipeline/filters/pipeline      [no test files]
?       pipeline/hash  [no test files]
?       pipeline/metrics       [no test files]
ok      pipeline/dispatcher    0.021s  coverage: 55.9% of statements
ok      pipeline/jobs  0.010s  coverage: 80.0% of statements
ok      pipeline/serial        0.004s  coverage: 65.3% of statements```
```

## 基准测试

**BenchmarkPipeline{队列个数}_{队列长度}-{运行的CPU个数}**

```bash
go  test -benchmem -run=^$ -bench . -benchtime=10x  -cpu=2,4,8   pipeline/dispatcher

goos: linux
goarch: amd64
pkg: pipeline/dispatcher
cpu: AMD EPYC 7K62 48-Core Processor
BenchmarkPipeline1_100-2              10           1931960 ns/op          338817 B/op      12578 allocs/op
BenchmarkPipeline1_100-4              10           1059881 ns/op          277638 B/op      10910 allocs/op
BenchmarkPipeline1_100-8              10           1405578 ns/op          277686 B/op      10869 allocs/op
BenchmarkPipeline100_100-2            10          17743815 ns/op         3132253 B/op     125071 allocs/op
BenchmarkPipeline100_100-4            10          15028784 ns/op         2934835 B/op     116644 allocs/op
BenchmarkPipeline100_100-8            10          15503481 ns/op         2852325 B/op     113000 allocs/op
BenchmarkPipeline1000_100-2           10         178007537 ns/op        32308539 B/op    1354064 allocs/op
BenchmarkPipeline1000_100-4           10         166615595 ns/op        29882658 B/op    1249756 allocs/op
BenchmarkPipeline1000_100-8           10         172102181 ns/op        28975368 B/op    1209757 allocs/op
BenchmarkPipeline10000_100-2          10        1967639871 ns/op        327559840 B/op  13800923 allocs/op
BenchmarkPipeline10000_100-4          10        1777576357 ns/op        305969252 B/op  12877903 allocs/op
BenchmarkPipeline10000_100-8          10        1663080771 ns/op        294666602 B/op  12385360 allocs/op
PASS
ok      pipeline/dispatcher    65.466s
```
