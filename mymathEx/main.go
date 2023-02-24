package main

import (
	"fmt"
)

func main() {
	fmt.Println("2+3=", mymath.Sum(2, 3))
	fmt.Println("4+7=", mymath.Sum(4, 7))
	fmt.Println("-2+-3=", mymath.Sum(-2, -3))
}
