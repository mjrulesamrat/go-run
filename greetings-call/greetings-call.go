package main

import (
	"fmt"
	"log"

	"greetings"
)

func main() {
	// set log message prifix and don't need to disable any flags
	log.SetPrefix("greetings: ")

	// names slice
	names := []string{"Jay", "Himani", "Liza"}

	// get the greeting message and print
	message, err := greetings.Hello("Jay")

	messages, errs := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	if errs != nil {
		log.Fatal(err)
	}

	// normal flow
	fmt.Println(message)
	fmt.Println(messages)
}
