package main

import "fmt"

func fibonacci() func() int {
	count := 0
	n1 := 0
	n2 := 1

	return func() int {
		if count==0 {
			count++
			return n1
		} else if count==1 {
			count++
			return n2
		}

		ret := n1+n2
		n1 = n2
		n2 = ret
		return ret
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
