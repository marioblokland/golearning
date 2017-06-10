package main

import "fmt"

func main() {
	channel := incrementor()
	channelSum := puller(channel)

	// Take the value off the channel and print it.
	// This also unblocks the main function which would
	// otherwise get stuck. since nothing would take the value
	// off the channelSum channel.
	fmt.Println(<-channelSum)
}

// This function will create a channel,
// iterate over the values 0 through 9 which will be taken
// by the puller function so the goroutine will be unblocked.
// Remember, if no other goroutine would take the value off the
// channel, the inner goroutine here would get stuck at 0
// since it would wait for the value to be taken off.
func incrementor() chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()

	return out
}

// This function has an inner goroutine which unblocks
// the goroutine of the incrementor function by pulling
// the values off its channel.
func puller(channel chan int) chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range channel {
			sum += n
		}
		out <- sum
		close(out)
	}()

	return out
}
