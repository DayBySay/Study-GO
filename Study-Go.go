package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (z float64) {
	z = 1.0
	for i := 0; i < 10; i++ {
		z = z - (z*z-x)/(2*z)
	}

	return
}

func main() {
	i := 1.0
	fmt.Println(Sqrt(i))
	fmt.Println(math.Sqrt(i))
}
