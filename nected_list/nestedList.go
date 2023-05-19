package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

var root = new(Node)

func addNode(t *Node, v int) int {
	if root == nil {
		t = &Node{v, nil}
		root = t
		return 0
	}
	if v == t.Value {
		fmt.Println("Node already exists:", v)
		return -1
	}
	if t.Next == nil {
		t.Next = &Node{v, nil}
		return -2
	}
	return addNode(t.Next, v)
}

func traverse(t *Node) {
	if t == nil {
		fmt.Println("-> Empty list!")
		return
	}
	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}

func lookupNode(t *Node, v int) bool {
	if root == nil {
		t = &Node{v, nil}
		root = t
		return false
	}
	if v == t.Value {
		return true
	}
	if t.Next == nil {
		return false
	}
	return lookupNode(t.Next, v)
}

func size(t *Node) int {
	if t == nil {
		fmt.Println("-> Empty list!")
		return 0
	}
	i := 0
	for t != nil {
		i++
		t = t.Next
	}
	return i
}

func remove(t *Node, v int) bool {
	if root == nil {
		fmt.Printf("Node with value: %d is not on the list", v)
		fmt.Println()
		return false
	}
	if root.Value == v {
		fmt.Printf("Node with value: %d is deleted", v)
		fmt.Println()
		root = root.Next
		return true
	}
	if t.Next == nil {
		fmt.Printf("Node with value: %d is not on the list", v)
		fmt.Println()
		return false
	}
	if t.Next.Value == v {
		if t.Next.Next == nil {
			t.Next = nil
		} else {
			t.Next = t.Next.Next
		}
		fmt.Printf("Node with value: %d is deleted\n", v)
		fmt.Println()
		return true
	}
	return remove(t.Next, v)
}

func main() {
	fmt.Println(root)
	root = nil
	traverse(root)
	addNode(root, 1)
	addNode(root, -1)
	traverse(root)
	addNode(root, 10)
	addNode(root, 5)
	addNode(root, 45)
	addNode(root, 5)
	addNode(root, 5)
	traverse(root)
	addNode(root, 100)
	traverse(root)
	if lookupNode(root, 100) {
		fmt.Println("Node exists!")
	} else {
		fmt.Println("Node does not exist!")
	}
	if lookupNode(root, -100) {
		fmt.Println("Node exists!")
	} else {
		fmt.Println("Node does not exist!")
	}
	remove(root, 10)
	traverse(root)
}
