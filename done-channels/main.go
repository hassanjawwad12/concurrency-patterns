package main

import (
	"fmt"
	"time"
)

// Start a goroutine which works by defualt but if the parent cancels it then it stops
// the done is being passed as a read-only channel
func doWork(done <-chan bool) {
	for {
		select {
		case <-done: // receive message from done channel
			fmt.Println("Done working")
			return
		default:
			fmt.Println("Waiting for data")
		}
	}
}

func main() {
	done := make(chan bool)
	go doWork(done)

	time.Sleep(1 * time.Second)
	// Close the channel to signal the goroutine to stop
	close(done)
}
