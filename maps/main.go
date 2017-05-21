package main

import (
	"fmt"
	"strings"
)

func main() {
	// A map with a key of int and a value of string
	myGreeting := map[int]string{
		0: "Guten Morgen!",
		1: "Good morning!",
		2: "Goeden morgen!",
	}

	for key, val := range myGreeting {
		fmt.Println(key, " - ", val)
	}

	fmt.Println(strings.Repeat("*", 40))

	anotherGreeting := make(map[int]string)
	anotherGreeting[0] = "Guten Morgen!"
	anotherGreeting[1] = "Good morning!"
	anotherGreeting[2] = "Goeden morgen!"

	for key, val := range anotherGreeting {
		fmt.Println(key, " - ", val)
	}

	// The following way of making a map should never be used
	var greeting map[int]string

	// This will cause an error, stating that we want to assign
	// a value to entry in a nil map. We can't add a value due to the
	// lack of an append function like we have in slices and because
	// a map is a reference type, this maps value is nil.
	greeting[0] = "Test"
}
