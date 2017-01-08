package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct{}

func (r MyReader) Read(str []byte) (int, error) {
	for i := 0; i < len(str); i++ {
		str[i] = 'A'
	}

	return len(str), nil
}

func main() {
	reader.Validate(MyReader{})
}
