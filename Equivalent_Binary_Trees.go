package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

func RecursiveWalk(t *tree.Tree, ch chan int) {
	if t==nil {
		return
	}

	RecursiveWalk(t.Left, ch)
	ch <- t.Value
	RecursiveWalk(t.Right, ch)
}

func Walk(t *tree.Tree, ch chan int) {
	if t==nil {
		return
	}

	RecursiveWalk(t.Left, ch)
	ch <- t.Value
	RecursiveWalk(t.Right, ch)
	close(ch)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if (v1!=v2) || (ok1!=ok2) {
			return false
		}
		if (ok1==false) || (ok2==false) {
			break
		}
	}

	return true
}

func main() {
	result := Same(tree.New(1), tree.New(2))
	fmt.Println("Are they same trees? :", result)
}
