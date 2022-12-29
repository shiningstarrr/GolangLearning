package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	x := []string{"hello", "world", "prabably", "not"}
	v, err := json.Marshal(x)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(v))
}
