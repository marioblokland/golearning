package main

import "fmt"

func main() {
	fmt.Println("Decimal \t Binary \t Hexadecimal \t UTF-8")
	for i := 0; i < 200; i++ {
		fmt.Printf("%d \t\t %b \t\t %x \t\t %q \n", i, i, i, i)
	}
}
