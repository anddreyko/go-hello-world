package main

import (
	"fmt"
	"log"

	"github.com/anddreyko/go-hello-world/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(greetings.Hello(message))
}
