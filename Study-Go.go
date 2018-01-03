package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (z float64) {
	z = 1.0
	for i := 0; i < 10; i++ {
		x := z - (z*z-x)/(2*z)
		if x == z {
			return
		}
		z = x
	}

	return
}

func main() {
	i := 5.0
	fmt.Println(Sqrt(i))
	fmt.Println(math.Sqrt(i))
}
