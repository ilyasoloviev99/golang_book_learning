package main

import (
	"sync"
	"sync/atomic"
	"time"
)

func main() {

}

func MutexCounter() int {
	goroutinesCounter := 0
	wg := sync.WaitGroup{}
	m := sync.Mutex{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			m.Lock()
			goroutinesCounter++
			m.Unlock()
			time.Sleep(time.Millisecond)
			m.Lock()
			goroutinesCounter--
			m.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	return goroutinesCounter
}

func AtomicCounter() int32 {
	goroutinesCounter := int32(0)
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&goroutinesCounter, 1)
			time.Sleep(time.Millisecond)
			atomic.AddInt32(&goroutinesCounter, -1)
			wg.Done()
		}()
	}

	wg.Wait()
	return goroutinesCounter
}
