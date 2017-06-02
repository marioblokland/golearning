package main

import (
	"fmt"
)

func main() {
	// This is a channel, over which we will
	// be able to communicate an integer.
	channel := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			// Put the current number to the channel
			channel <- i
			// The code stops here, until something
			// takes the value off the channel.
		}
		close(channel)
	}()

	// Since we receive a value from the channel
	// the above function can go and execute its next
	// iteration.
	for n := range channel {
		fmt.Println(n)
	}
}
