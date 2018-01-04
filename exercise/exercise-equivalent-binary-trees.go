package exercise

import (
	"golang.org/x/tour/tree"
	"reflect"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	_Walk(t, ch)
	close(ch)
}

func _Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	_Walk(t.Left, ch)
	ch <- t.Value
	_Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	s1 := make([]int, 0)
	s2 := make([]int, 0)

	c1 := make(chan int)
	c2 := make(chan int)

	go Walk(t1, c1)
	go Walk(t2, c2)

	for v := range c1 {
		s1 = append(s1, v)
	}

	for v := range c2 {
		s2 = append(s2, v)
	}

	return reflect.DeepEqual(s1, s2)
}
