package exercise

import (
	"fmt"
	"golang.org/x/tour/tree"
	"testing"
)

func TestSame_true(t *testing.T) {
	if !Same(tree.New(1), tree.New(1)) {
		t.Fatal("Same(tree.New(1), tree.New(1)) should be true")
	}
}

func TestSame_false(t *testing.T) {
	if Same(tree.New(2), tree.New(1)) {
		t.Fatal("Same(tree.New(2), tree.New(1)) should be false")
	}
}

func ExampleWalk() {
	c := make(chan int)
	go Walk(tree.New(1), c)
	for i := range c {
		fmt.Println(i)
	}
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
	// 7
	// 8
	// 9
	// 10
}
