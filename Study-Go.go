package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) (p [][]uint8) {
	r := make([][]uint8, dy)
	for i := range r {
		r[i] = make([]uint8, dx)
		for j := range r[i] {
			r[i][j] = uint8(i ^ j + (i+j)/2)
		}
	}

	p = r
	return
}

func main() {
	pic.Show(Pic)
}
