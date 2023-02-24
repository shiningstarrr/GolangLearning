package main

import (
	"PracticeCode/01/mymath"
	"fmt"
)

func main() {
	fmt.Println("2+3 = ", mymath.Sum(2, 3))
	fmt.Println("10+10 = ", mymath.Sum(10, 10))
	fmt.Println("6+9 = ", mymath.Sum(6+9))
}
