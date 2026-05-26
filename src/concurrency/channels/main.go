package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		println("Hello from a goroutine!")
		ch <- 42 // Signal that the goroutine is done
	}()
	fmt.Println(<-ch)
	close(ch) // Close the channel when done
	// fmt.Println("Hello, World!")
}
