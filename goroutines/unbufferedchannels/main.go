package main

import (
	"fmt"
	"time"
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
	}()

	go func() {
		for {
			// The number will be taken off the channel and
			// the above function will execute further.
			fmt.Println(<-channel)
		}
	}()

	// Dont end the script directly, wait a second first,
	// so that the above concurrent functions will be finished.
	// Without this statement, one would see nothing in the terminal.
	time.Sleep(time.Second)
}
