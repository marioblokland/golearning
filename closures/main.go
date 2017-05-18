package main

import (
	"fmt"
)

func wrapper() func() int {
	var x int // x will be initialized automatically with int zero value: 0
	return func() int {
		x++
		return x
	}
}

func main() {
	increment := wrapper()
	fmt.Println(increment()) // will print 1
	fmt.Println(increment()) // will print 2
}

/*
Closure helps us to limit the scope of variables used by multiple functions.
Without closure, for two or more funcs to have access to the same variable,
that variable would need to be at package scope
*/
