package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func repeatFunc[T any, K any](done <-chan K, fn func() T) <-chan T {
	// send operation on an unbuffered channel blocks the sending go-routine until
	//the value is read by the receiving go-routine
	stream := make(chan T)

	go func() {
		defer close(stream)
		for {
			select {
			case <-done:
				return
			case stream <- fn():
			}
		}
	}()
	return stream
}

func take[T any, K any](done <-chan K, stream <-chan T, n int) <-chan T {
	takeStream := make(chan T)

	go func() {
		defer close(takeStream)
		for range n {
			select {
			case <-done:
				return
			case takeStream <- <-stream: // write a value from stream to takeStream
				// <-stream means reading a value from stream channel
			}
		}
	}()
	return takeStream
}

func primeFinder(done <-chan int, stream <-chan int) <-chan int {

	isPrime := func(num int) bool {
		for i := num - 1; i > 1; i-- {
			if num%i == 0 {
				return false
			}
		}
		return true
	}

	// make another stream to write data on it
	primeStream := make(chan int)

	go func() {
		defer close(primeStream)
		for {
			select {
			case <-done:
				return
			case num := <-stream:
				if isPrime(num) {
					primeStream <- num
				}
			}
		}
	}()
	return primeStream
}

func fanIn[T any](done <-chan int, channels ...<-chan T) <-chan T {
	var wg sync.WaitGroup
	fanInStream := make(chan T)

	//read values from all the channels and write to the fanInStream
	transfer := func(ch <-chan T) {
		defer wg.Done()
		for val := range ch {
			select {
			case <-done:
				return
			case fanInStream <- val:
			}
		}
	}

	// pass dafa from the concurrently running  primefinder channels to the fanInStream
	for _, ch := range channels {
		wg.Add(1)
		go transfer(ch)
	}

	// wait is finished closed the stream
	go func() {
		wg.Wait()
		close(fanInStream)
	}()
	return fanInStream
}

func main() {
	start := time.Now()
	done := make(chan int)
	defer close(done)

	randomNumber := func() int {
		return rand.Intn(50000000)
	}

	randIntStream := repeatFunc(done, randomNumber)

	// fan out
	CPUCount := runtime.NumCPU()
	primtFinderChannels := make([]<-chan int, CPUCount) // slice of channels
	for i := range CPUCount {
		// function called for each CPU and result is stored in the channel slice
		primtFinderChannels[i] = primeFinder(done, randIntStream)
	}

	// fan in
	// merge all the channels into one channel
	fannedInStream := fanIn(done, primtFinderChannels...)
	for rand := range take(done, fannedInStream, 10) {
		println(rand)
	}

	fmt.Println(time.Since(start))
}
