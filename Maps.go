package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	words := strings.Fields(s)

	for i := 0; i < len(words); i++ {
		v, ok := m[words[i]]
		if ok {
			m[words[i]] = v + 1
		} else {
			m[words[i]] = 1
		}
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
