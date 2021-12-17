package main

import (
	"fmt"

	"github.com/mvsrgc/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
