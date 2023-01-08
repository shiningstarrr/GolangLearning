package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")
	val := 3
	p := &val
	fmt.Println("address", *p)
	double(&val)
	fmt.Println(val)
	fmt.Println(*p)
}

func double(x *int) {
	*x = *x * 2

}
