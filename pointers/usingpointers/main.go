package main

import "fmt"

func main() {
	a := 43

	fmt.Println("Value of a: ", a)
	fmt.Println("Memory address of a: ", &a)

	// Normally the shorthand notation b := &a should be used instead,
	// which ends up in the same result, but the longer version makes
	// understanding the code easier.
	var b *int = &a // make b a pointer to the address where the int of a is stored
	fmt.Println("Value of b is the memory address of a: ", b)

	// Dereference b with the * sign before the variable name
	fmt.Println("Value of b after dereferencing: ", *b)

	// Change the value at the memory address of a and the referenced b
	*b = 42

	fmt.Println("Value of a, after changing the value at the memory address: ", a)
}
