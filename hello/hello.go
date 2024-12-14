package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	// fmt.Println(quote.Go())
	message := greetings.Hello("Sudhendra")
	fmt.Println(message)
}
