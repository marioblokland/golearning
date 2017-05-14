package main

import "fmt"

const metersToYards float64 = 1.09361

func main() {
	var meters float64
	fmt.Println("Enter meters swam: ")
	fmt.Scan(&meters) // variable has to be provided with the ampersand sign
	yards := meters * metersToYards
	fmt.Println(meters, "meters is", yards, "yards.")
}
