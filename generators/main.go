package main

import (
	"fmt"
	"math/rand"
	"time"
)

func repeatFunc[T any, K any](done <-chan K, fn func() T) <-chan T {
	// Create a channel of type T
	stream := make(chan T)

	// this go routine continue tom put data on the stream
	go func() {
		// when go routine ends close the stream
		defer close(stream)
		for {
			select {
			case <-done:
				return
				// Write the result to call to the function to the stream (returns type T which is put on the channel)
			case stream <- fn(): // it calls this function if the done channel is not closed which it checks on each iteration
			}
		}
	}()
	return stream
}

func main() {
	done := make(chan struct{}) // Use empty struct as a signal-only channel.

	randomNumber := func() int {
		return rand.Intn(500)
	}
	go func() {
		time.Sleep(1 * time.Second)
		close(done) // Signal to stop
	}()

	for rand := range repeatFunc(done, randomNumber) {
		println(rand)
	}
	fmt.Println("Stopped generating random numbers")
}
