package main

import (
	"context"
	"fmt"
	"sync/atomic"
)

func main() {
	value := 0
	value++
	fmt.Println(value)

	value2 := int32(0)
	atomic.AddInt32(&value2, 1)
	fmt.Println(value2)

	value3 := atomic.Value{}
	value3.Swap(100)
	result := value3.Load()
	fmt.Println(result)

	ctx := context.Background()
	fmt.Println(ctx)
}
