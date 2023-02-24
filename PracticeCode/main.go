package main

import (
	"errors"
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		e := errors.New("more coffee needed")
		return f, sqrtError{"50", "99", e}
	}
	return 42, nil
}

func main() {
	fmt.Println("start")
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}

}
