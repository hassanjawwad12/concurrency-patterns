package main

import (
	"fmt"
	"math/rand"
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

func main() {
	start := time.Now()
	done := make(chan int)
	defer close(done)

	randomNumber := func() int {
		return rand.Intn(500)
	}

	// runs infinitely to produce random numbers
	randIntStream := repeatFunc(done, randomNumber)

	// pass the random number stream to primeFinder
	// primeFinder will filter out the prime numbers from the stream
	prime := primeFinder(done, randIntStream)

	// take the first 10 prime numbers from the prime stream and stops the repeatFunc
	for rand := range take(done, prime, 10) {
		println(rand)
	}
	fmt.Println(time.Since(start))
}
