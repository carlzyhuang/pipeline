package dispatcher

import (
	"fmt"
	"math"
	"strconv"
	"sync"
	"testing"
	"time"

	"pipeline/jobs"
	"pipeline/serial"
)

func TestPipeline(t *testing.T) {
	cases := []struct {
		Name      string
		HashKey   string
		PostData  int
		LoopCount int
		Expected  int
	}{
		{
			"hash key 1",
			"1",
			100,
			100,
			200,
		},
		{
			"hash key max",
			strconv.Itoa(math.MaxUint32),
			10000,
			2000,
			12000,
		},
	}

	dispatcher := NewDispatcher(&serial.DefaultSerializer[string]{}, jobs.NewWorkQueue(jobs.GetDefaultConfig()))
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			wg := sync.WaitGroup{}

			for i := 0; i < c.LoopCount; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()

					dispatcher.Post(c.HashKey, func() {
						c.PostData++
					})
				}()
			}

			wg.Wait()

			for {
				buffLen, _ := dispatcher.GetJobsBuffLen(c.HashKey)
				if buffLen == 0 {
					break
				}
				time.Sleep(time.Millisecond)
			}

			if c.Expected != c.PostData {
				t.Fatalf("%s: expected %v, got %v", c.Name, c.Expected, c.PostData)
			}
		})

	}
}

func TestPipelineFIFO(t *testing.T) {

	dispatcher := NewDispatcher(&serial.DefaultSerializer[string]{}, jobs.NewWorkQueue(jobs.GetDefaultConfig()))

	count := 0
	dispatcher.Post("1", func() {
		if count != 0 {
			t.Fatalf("%s: expected %v, got %v", "FIFO test error", 0, count)
		}
		count++
	})

	dispatcher.Post("1", func() {
		if count != 1 {
			t.Fatalf("%s: expected %v, got %v", "FIFO test error", 1, count)
		}
		count++
	})

	dispatcher.Post("1", func() {
		if count != 2 {
			t.Fatalf("%s: expected %v, got %v", "FIFO test error", 2, count)
		}
		count++
	})

	time.Sleep(time.Millisecond * 10)
	if count != 3 {
		t.Fatalf("%s: expected %v, got %v", "FIFO test error", 3, count)
	}
}

// func TestRelaunch(t *testing.T) {
// 	cases := []struct {
// 		Name string
// 		Cfg  *jobs.PipelineConfig
// 	}{
// 		{
// 			"case 1",
// 			&jobs.PipelineConfig{
// 				MaxWorkerQueueCount: 10000,
// 				MaxJobsPerWorker:    1,
// 			},
// 		},
// 		{
// 			"case 2",
// 			&jobs.PipelineConfig{
// 				MaxWorkerQueueCount: 1,
// 				MaxJobsPerWorker:    1,
// 			},
// 		},
// 	}

// 	for _, c := range cases {
// 		t.Run(c.Name, func(t *testing.T) {
// 			err := RelaunchDefaultWorkerQueue(c.Cfg)
// 			if err != nil {
// 				t.Fatalf("%s: relaunch err %v", c.Name, err)
// 			}
// 		})
// 	}

// }

func benchmarkPipeline(b *testing.B, keyCount, queueCount int) {
	type cases struct {
		Name     string
		HashKey  uint64
		PostData int
		Expected int
	}

	dispatcher := NewDispatcher(&serial.DefaultSerializer[uint64]{}, jobs.NewWorkQueue(jobs.GetDefaultConfig()))
	for n := 0; n < b.N; n++ {
		for i := 0; i < keyCount; i++ {
			wg := sync.WaitGroup{}

			c := &cases{
				Name:     fmt.Sprintf("case %d", i),
				HashKey:  uint64(i),
				PostData: 0,
				Expected: queueCount,
			}

			for i := 0; i < queueCount; i++ {
				wg.Add(1)
				go func() {
					dispatcher.Post(c.HashKey, func() {
						defer wg.Done()

						c.PostData++
					})
				}()
			}

			wg.Wait()

			if c.Expected != c.PostData {
				b.Fatalf("%s: %d expected %d, got %d", c.Name, c.HashKey, c.Expected, c.PostData)
			}
		}

	}
}

// key 1 queue 1000
func BenchmarkPipeline1_100(b *testing.B) { benchmarkPipeline(b, 1, 1000) }

func BenchmarkPipeline100_100(b *testing.B) { benchmarkPipeline(b, 100, 100) }

func BenchmarkPipeline1000_100(b *testing.B) { benchmarkPipeline(b, 1000, 100) }

func BenchmarkPipeline10000_100(b *testing.B) { benchmarkPipeline(b, 10000, 100) }
