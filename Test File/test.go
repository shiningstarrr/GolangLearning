package main

import (
	"fmt"
)

func init() {
	fmt.Println("start")
}

func main() {
	m := map[string](int){
		"first":  1,
		"second": 2,
		"third":  3,
	}
	if _, ok := m["second"]; !ok {
		fmt.Println("not exist")
	} else {
		fmt.Println("founded")
	}

}
