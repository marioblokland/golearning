package main

import "fmt"

const (
	// iota always increments the next iota variable value by 1.
	// So the first iota is 0, the next would be 0 + 1 and so on.
	_ = iota // 0 and "_" says, that this variable is not used
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	MB = 1 << (iota * 10) // 1 << (2 * 10)
	GB = 1 << (iota * 10) // 1 << (3 * 10)
)

func main() {
	fmt.Println("Binary \t\t\t\t\t Decimal")
	fmt.Printf("%b \t\t\t\t %d \n", KB, KB)
	fmt.Printf("%b \t\t\t %d \n", MB, MB)
	fmt.Printf("%b \t %d \n", GB, GB)
}
