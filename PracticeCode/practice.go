package main

import (
	"fmt"
	"sort"
)

func main() {

	people := []struct {
		Name string
		Age  int
	}{
		{"Alice", 25},
		{"Zlizabeth", 75},
		{"Alice", 75},
		{"Gob", 75},
		{"Alice", 75},
		{"Fob", 25},
		{"Colin", 25},
		{"Elizabeth", 25},
	}
	fmt.Println("before sort: ", people)
	sort.Slice(people, func(i, j int) bool { return people[i].Name < people[j].Name })
	fmt.Println("after sort: ", people)

}
