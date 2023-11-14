package main

import (
	"fmt"
	"log"

	"github.com/anddreyko/go-hello-world/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Drew", "Alex", "Paddington"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
