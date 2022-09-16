package hello

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
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
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}
func NewMethod(name string) string {
	return name
} // init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}

/*
In this code, you:

For Randam section

Add a randomFormat function that returns a randomly selected format for a greeting message. Note that randomFormat starts with a lowercase letter, making it accessible only to code in its own package (in other words, it's not exported).
In randomFormat, declare a formats slice with three message formats. When declaring a slice, you omit its size in the brackets, like this: []string. This tells Go that the size of the array underlying the slice can be dynamically changed.
Use the math/rand package to generate a random number for selecting an item from the slice.
Add an init function to seed the rand package with the current time. Go executes init functions automatically at program startup, after global variables have been initialized. For more about init functions, see Effective Go.
In Hello, call the randomFormat function to get a format for the message you'll return, then use the format and name value together to create the message.
Return the message (or an error) as you did before.

*/
