package hello

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(Hello("Karthik"))
	return
}

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
func NewMethod(name string) string {
	return name
}
