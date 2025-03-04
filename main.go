package main

import (
	"fmt"
	"time"
)

func random(num string) {
	fmt.Println("Random number is: ", num)
}

func main() {
	// go routines independent of each other
	go random("123")

	time.Sleep(2 * time.Second)
	fmt.Println("Hello, World!")
}
