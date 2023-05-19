package main

import (
	"fmt"
	"sort"
)

type Interface interface {
	sort.Interface
	Push(x interface{}) // вставляет x как элемент Len()
	Pop() interface{}   // удаляет и возвращает элемент Len() - 1
}

type heapFloat32 []float32

func (n *heapFloat32) Pop() interface{} {
	old := *n
	x := old[len(old)-1]
	new := old[0 : len(old)-1]
	*n = new
	return x
}

func (n *heapFloat32) Push(x interface{}) {
	*n = append(*n, x.(float32))
}

type Point struct {
	X, Y int
}

func main() {
	var p Point

	if p == struct{ X, Y int }{} {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
