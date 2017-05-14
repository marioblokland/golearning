package main

import "fmt"

func main() {
	a := 43

	fmt.Println("Value of a: ", a)
	fmt.Println("Memory address of a: ", &a)

	// Normally the shorthand notation b := &a should be used instead,
	// which ends up in the same result, but the longer version makes
	// understanding the code easier.
	// b is of type "int pointer" --  the * is part of the type
	var b *int = &a // make b a pointer to the address where the int of a is stored
	fmt.Println("Value of b is the memory address of a: ", b)

	// Dereference b with the * sign before the variable name
	fmt.Println("Value of b after dereferencing: ", *b)

	// b is an int pointer
	// b points to the memory address where an int is stored
	// to see the value in that memory address, put * before the variable
	// which is known as dereferencing.
}
