package main

import (
	"fmt"
	"greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Something from greetings module")
	fmt.Println(message)
}
