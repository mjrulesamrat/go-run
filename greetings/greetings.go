package greetings

import "fmt"

// Hello returns a greetings with given name
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v Welcome to go-world!", name)
	return message
}
