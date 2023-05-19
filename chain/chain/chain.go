package main

import (
	"fmt"
	"sync"
	"time"
)

var ch = make(chan string)

func main() {
	go someMessages()

	time.Sleep(time.Second)

	for i := range ch {
		fmt.Println(i)
	}
}

func someMessages() {
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			ch <- fmt.Sprintf("Hello - %v", i)
			wg.Done()
		}(i)
	}

	wg.Wait()
	close(ch)
}
