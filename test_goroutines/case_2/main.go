package main

import (
	"fmt"
	"sync"
)

func main() {
	//runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup

	for x := 0; x < 10; x++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			fmt.Printf("goroutine: %v\n", x)
		}(x)
	}

	wg.Wait()
	fmt.Println("Done!")
}
