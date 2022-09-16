package hello

import (
	"errors"
	"fmt"
)

// func main() {
// 	fmt.Println("Hello, World!")
// 	fmt.Println(Hello("Karthik"))
// 	return
// }

/*
In this code, you: For error section

Change the function so that it returns two values: a string and an error. Your caller will check the second value to see if an error occurred. (Any Go function can return multiple values. For more, see Effective Go.)
Import the Go standard library errors package so you can use its errors.New function.
Add an if statement to check for an invalid request (an empty string where the name should be) and return an error if the request is invalid. The errors.New function returns an error with your message inside.
Add nil (meaning no error) as a second value in the successful return. That way, the caller can see that the function succeeded.

*/

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
func NewMethod(name string) string {
	return name
}
