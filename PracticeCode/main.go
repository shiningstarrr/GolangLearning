package main

import (
	"fmt"
	"strings"
)

func removeProfanity(message *string) {
	replacer := strings.NewReplacer("dang", "****", "shoot", "****", "heck", "****")
	s := replacer.Replace(*message)
	*message = s

}

// don't touch below this line

func main() {
	messages := []string{
		"well shoot, this is awful",
		"dang robots",
		"dang them to heck",
	}

	for _, message := range messages {
		removeProfanity(&message)
		fmt.Println(message)
	}
}
