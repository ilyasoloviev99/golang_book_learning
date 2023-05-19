package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	cs := map[string]int{"касса": 0, "касса2": 0}
	mu := &sync.Mutex{}

	for i := 0; i < 100; i++ {
		go func(k int) {
			mu.Lock()
			defer mu.Unlock()
			cs["касса"] += 1
		}(i)
	}

	for i := 0; i < 100; i++ {
		go func(k int) {
			mu.Lock()
			defer mu.Unlock()
			cs["касса2"] += 1
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println(cs)
}
