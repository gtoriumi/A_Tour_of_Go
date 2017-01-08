package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0

	for {
		z2 := z - (z*z-2)/(2*z)

		if math.Abs(z2-z) < 0.000001 {
			return z2
		}

		z = z2
	}
}

func main() {
	fmt.Println(Sqrt(2))
}
