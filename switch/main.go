package main

import (
	"fmt"
)

// In Go, switch statemens don't need a break statement.
// If one wants to keep going on a specific case, you
// have to declare it explicitly with a fallthrough statement.
func main() {
	var name string
	fmt.Println("Enter your name: ")
	fmt.Scan(&name)

	switch name {
	case "Kaylem":
		fmt.Println("Wassup Kaylem")
		fallthrough // The next case will also be executed
	case "Andrea":
		fmt.Println("Wassup Andrea")
		fallthrough // The next case will also be executed
	case "Mario":
		fmt.Println("Wassup Mario")
	case "Denise", "Claudia", "Malick": // You can also handle multiple cases at once
		fmt.Println("Wassup Denise, Claudia or Malick")
	default:
		fmt.Println("Wassup Unknown")
	}
}
