package main

import (
	"fmt"
)

// This is a function, which states that it
// returns another function which returns a string
func makeGreeter() func() string {

	// Function expression
	myGreeter := func() string {
		return "Hello, friend!"
	}

	return myGreeter
}

// This is a function, which states that it
// returns another function which returns a string
func makeAnotherGreeter() func() string {
	return func() string {
		return "Hello, friend 2.0!"
	}
}

func main() {
	greet := makeGreeter()
	anotherGreet := makeAnotherGreeter()
	fmt.Println(greet())
	fmt.Println(anotherGreet())
}
