package main

import (
	"fmt"
	"example.com/greetings"
)

import "rsc.io/quote"

func main() {
	fmt.Println("Hello, world")
	fmt.Println(quote.Go())
	message := greetings.Hello("Deema")
	fmt.Println(message)
}
