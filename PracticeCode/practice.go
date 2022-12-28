package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	fmt.Println("start")
	s := `pass123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != err {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(bs)
}
