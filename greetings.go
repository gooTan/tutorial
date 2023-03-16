package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hi %c. Welcome!", name)
	return message
}
