package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	var fName string
	var lName string
	fmt.Println("enter first name ")
	fmt.Scanln(&fName)

	fmt.Println("enter last name: ")
	fmt.Scanln(&lName)
	fmt.Println("full name is: " + fName + " " + lName)
	s := `pass123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	fmt.Println(bs)
	fmt.Println(err)

}
