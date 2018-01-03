package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	c := 0
	x := 0
	y := 1
	return func() int {
		c++
		if c == 1 {
			return 0
		} else if c == 2 {
			return 1
		}

		r := x + y
		x = y
		y = r
		return r
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
