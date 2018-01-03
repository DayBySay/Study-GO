package calc

import (
    "testing"
    "fmt"
)

func TestSum (t *testing.T) {
    if sum(1, 2) != 3 {
        t.Fatal("sum(1,2) should be 3, but doesn't match")
    }
}

func ExampleSum () {
    fmt.Println(sum(1, 3))
    // OutPut: 4
}
