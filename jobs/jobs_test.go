package jobs

import (
	"math"
	"sync"
	"testing"
	"time"
)

func TestWorkerQueue(t *testing.T) {
	cases := []struct {
		Name      string
		HashKey   uint64
		PostData  int
		LoopCount int
		Expected  int
	}{
		{
			"hash key 1",
			1,
			100,
			100,
			200,
		},
		{
			"hash key max",
			math.MaxUint64,
			1000,
			200,
			1200,
		},
	}

	defaultWorkQueue := NewWorkQueue(GetDefaultConfig())

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			wg := sync.WaitGroup{}

			for i := 0; i < c.LoopCount; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()

					defaultWorkQueue.Dispatch(c.HashKey, func() {
						c.PostData++
					})
				}()
			}

			wg.Wait()

			for {
				if defaultWorkQueue.JobsBuffLen(c.HashKey) == 0 {
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

func TestJobPanic(t *testing.T) {
	cases := []struct {
		Name      string
		HashKey   uint64
		PostData  int
		LoopCount int
		Expected  int
	}{
		{
			"panic case 1",
			1,
			10,
			10,
			15,
		},
	}

	defaultWorkQueue := NewWorkQueue(GetDefaultConfig())

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			wg := sync.WaitGroup{}
			for i := 0; i < c.LoopCount; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()

					defaultWorkQueue.Dispatch(c.HashKey, func() {
						if c.PostData == c.Expected {
							panic("this is a test panic")
						}
						c.PostData++
					})
				}()
			}

			wg.Wait()

			for {
				if defaultWorkQueue.JobsBuffLen(c.HashKey) == 0 {
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
