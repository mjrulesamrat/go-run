package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greetings with given name
func Hello(name string) (string, error) {
	// handle empty name case
	if name == "" {
		return "", errors.New("name this man")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// init sets initial values for variables used in the function
func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	// a slice of message formats
	formats := []string{
		"Hi, %v Welcome!",
		"Great to see you %v",
		"Welcome to Go world, %v",
	}

	// return a randomly selected message format by specifying
	// a random index for the slice of formats
	return formats[rand.Intn(len(formats))]
}
