package main

import "fmt"

// slice of int to a channel of int
func sliceToChannel(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n // send n to the channel
		}
		close(out) // close the channel when loop is done
	}()
	return out
}

// read only channel
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n // send the square of n to the channel
		}
		close(out) // close the channel when loop is done
	}()
	return out
}

func main() {
	// pipeline will be orchestrated via our main function

	//input
	nums := []int{1, 2, 3, 4, 5}
	// stage 1
	dataChannel := sliceToChannel(nums)

	// stage 2 : pass the output from stage1 to another function
	finalChannel := square(dataChannel)

	// stage 3 : output result of entire pipeline
	for n := range finalChannel {
		fmt.Println(n)
	}
}
