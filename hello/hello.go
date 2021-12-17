package main

import (
	"fmt"
	"github.com/mvsrgc/greetings"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Gladys")

	if err != nil {
		log.Fatal(err)
	}

	log.Print(message)
	fmt.Println(message)
}
