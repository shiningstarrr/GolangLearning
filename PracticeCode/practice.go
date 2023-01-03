package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{5, 8, 2, 43, 293, 88, 66, 51, 78}
	xs := []string{"random", "rainbor", "delight", "in", "tornnaod", "summer", "under"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println("after sort:", xi)

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println("after sort:", xs)

	x := sort.Search(len(xi), func(i int) bool { return xi[i] == 88 })
	fmt.Println("x equals to", x)

}
