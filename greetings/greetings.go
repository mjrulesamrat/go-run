package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greetings with given name
func Hello(name string) (string, error) {
	// handle empty name case
	if name == "" {
		return "", errors.New("name this man")
	}
	message := fmt.Sprintf("Hi, %v Welcome to go-world!", name)
	return message, nil
}
