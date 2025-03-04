package main

import "fmt"

func main() {
	// channel works as FIFO

	myChannel := make(chan string)
	anotherChannel := make(chan string)

	go func() {
		// writing data to the channel
		myChannel <- "data"
	}()

	go func() {
		// writing data to the channel
		anotherChannel <- "cow"
	}()

	// msg := <-myChannel // this myChannel is blocking
	// fmt.Println(msg)

	// main function is reading the data from the channel
	// this is where the forked go routine is joining back the main go routine
	select {
	case msgFromMyChannel := <-myChannel:
		fmt.Println(msgFromMyChannel)
	case msgFromAnotherChannel := <-anotherChannel:
		fmt.Println(msgFromAnotherChannel)
	}
}
