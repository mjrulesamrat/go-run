package main

import (
	"fmt"

	"greetings"
)

func main() {
	// get the greeting message and print
	message := greetings.Hello("Jay")
	fmt.Println(message)
}
