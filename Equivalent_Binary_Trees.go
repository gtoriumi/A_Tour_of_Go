package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

func Walk(t *tree.Tree, ch chan int) {
	if t==nil {
		return
	}

	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i:=0; i<10; i++ {
		if <-ch1 != <-ch2 {
			return false
		}
	}

	return true
}

func main() {
	result := Same(tree.New(5), tree.New(2))
	fmt.Println("Are they same trees? :", result)
}
