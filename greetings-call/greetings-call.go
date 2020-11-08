package main

import (
	"fmt"
	"log"

	"greetings"
)

func main() {
	// set log message prifix and don't need to disable any flags
	log.SetPrefix("greetings: ")

	// get the greeting message and print
	message, err := greetings.Hello("Jay")

	if err != nil {
		log.Fatal(err)
	}
	// normal flow
	fmt.Println(message)
}
