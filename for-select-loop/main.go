package main

func main() {

	// buffered channel with capacity 3
	// for asynchronous communication we need to have a buffered channel
	charChannel := make(chan string, 3)
	chars := []string{"a", "b", "c", "d", "e"}

	// send data to channel
	for _, char := range chars {
		select {
		case charChannel <- char:
			println("Sent:", char)
		default:
			println("Channel is full. Cannot send:", char)
		}
	}
	// after we iterate all the characters and put them in the channel, we close the channel
	close(charChannel)

	// we loop over a closed channel and still receive the residual data that was put onto it
	for char := range charChannel {
		println("Received:", char)
	}
}
