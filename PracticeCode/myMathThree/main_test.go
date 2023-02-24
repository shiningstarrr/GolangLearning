package mymaththree

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}
	tests := []test{test{[]int{21, 21}, 42}, test{[]int{3, 4, 5}, 12}, test{[]int{1, 2}, 3}, test{[]int{-1, -2}, -3}}

	for _, v := range tests {
		x := Sum(v.data...)
		if x != v.answer {
			t.Error("Expected", v.answer, "Got", x)

		}
	}
}
func ExampleSum() {
	fmt.Println(Sum(2, 3))
}
