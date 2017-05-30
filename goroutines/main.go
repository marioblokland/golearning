package main

import (
	"fmt"
)

// When the "go" keyword is used, this program
// will execute concurrent and parallel. Before Go 1.5,
// one had to explicitly set GOMAXPROCS to the maximum
// available CPUs in order to get parallelism.
func main() {
	// Without the keyword "go", these
	// functions will run procedural, this after
	// each other. On my machine, the main function
	// will take about between 0.5s and 0.6s
	// in order to be finished.
	//foo()
	//bar()

	// With the "go" keyword, the functions will
	// execute concurrent and take about between 0.15s
	// and 0.18s on my machine.
	go foo()
	go bar()
}

func foo() {
	for i := 0; i < 100000; i++ {
		fmt.Println("Foo:", i)
	}
}

func bar() {
	for i := 0; i < 100000; i++ {
		fmt.Println("Foo:", i)
	}
}
