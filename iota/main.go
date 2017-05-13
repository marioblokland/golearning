package main

import "fmt"

const (
	// iota always increments the next iota variable value by 1.
	// So the first iota is 0, the next would be 0 + 1 and so on.
	// In contrast to other languages, constants in go should not
	// be written in uppercase.
	_  = iota             // 0 and "_" says, that this variable is not used
	kb = 1 << (iota * 10) // 1 << (1 * 10)
	mb = 1 << (iota * 10) // 1 << (2 * 10)
	gb = 1 << (iota * 10) // 1 << (3 * 10)
)

func main() {
	fmt.Println("Binary \t\t\t\t\t Decimal")
	fmt.Printf("%b \t\t\t\t %d \n", kb, kb)
	fmt.Printf("%b \t\t\t %d \n", mb, mb)
	fmt.Printf("%b \t %d \n", gb, gb)
}
