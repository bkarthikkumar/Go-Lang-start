package newModule

import (
	"fmt"

	helloModule "example/hello"
)

func main() {
	// Get a greeting message and print it.
	message := helloModule.Hello("kingName")
	fmt.Println(message)
	fmt.Println(helloModule.NewMethod("KK"))
}
