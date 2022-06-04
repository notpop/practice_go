package lib

import (
	"fmt"
	"testing"
)

func Example() {
	v := Avarage([]int{1, 2, 3, 4, 5})
	fmt.Println(v)
}

// func ExampleAverage() {
// 	v := Avarage([]int{1, 2, 3, 4, 5})
// 	fmt.Println(v)
// }

func TestAvarage(t *testing.T) {
	v := Avarage([]int{1, 2, 3, 4, 5})
	if v != 3 {
		t.Error("Expected got 3", v)
	}
}
