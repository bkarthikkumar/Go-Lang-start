package main

import (
	"fmt"
	"log"

	helloModule "example/hello"
)

/*
In this code, you:

Configure the log package to print the command name ("greetings: ") at the start of its log messages, without a time stamp or source file information.
Assign both of the Hello return values, including the error, to variables.
Change the Hello argument from Gladys’s name to an empty string, so you can try out your error-handling code.
Look for a non-nil error value. There's no sense continuing in this case.
Use the functions in the standard library's log package to output error information. If you get an error, you use the log package's Fatal function to print the error and stop the program.
*/

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Request greeting messages for the names.
	messages, err := helloModule.Hellos(names)

	// Get a greeting message and print it.
	// Request a greeting message.
	message, err := helloModule.Hello("Karthik")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned map of
	// messages to the console.
	fmt.Println(messages)
	fmt.Println(message)
	fmt.Println(helloModule.NewMethod("KK"))
}

/*
With these changes, you:

Create a names variable as a slice type holding three names.
Pass the names variable as the argument to the Hellos function.
*/
