package main

import "fmt"

func main() {
	var fName string
	var lName string
	fmt.Println("enter first name ")
	fmt.Scanln(&fName)

	fmt.Println("enter last name: ")
	fmt.Scanln(&lName)
	fmt.Println("full name is: " + fName + " " + lName)

}
