package main

import (
	"fmt"
	"strings"
)

func main() {
	// This first slice has a length of 3 and a capacity of 3, since we initialized it
	// with 3 values. If the capacity gets exceeded, a new slice will be created with
	// the capacity of 4 and so on. This could cost performance issues at
	// some point
	fmt.Println("Slice made with shorthand notation")
	someSlice := []string{"Good morning", "Guten Morgen", "Goeden dag"}
	fmt.Println(someSlice)
	fmt.Println(len(someSlice))
	fmt.Println(cap(someSlice))

	fmt.Println(strings.Repeat("=", 40))

	// This slice has a current length of 3 and a capacity of 10.
	// We can add new entries to the slice and the capacity will still be 10.
	// If we exceed 10, the capacity will be doubled to 20 and so on.
	anotherSlice := make([]int, 3, 10)
	fmt.Println("Slice made with make")
	fmt.Println(anotherSlice)
	fmt.Println(len(anotherSlice))
	fmt.Println(cap(anotherSlice))
}
