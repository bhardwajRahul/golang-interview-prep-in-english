package main

import "fmt"

func main() {
	// Goroutines are lightweight threads managed by the Go runtime.
	go func() {
		println("Hello from a goroutine!")
	}()
	// Wait for the goroutine to finish (in a real application, use sync.WaitGroup or channels)
	select {
	default:
		fmt.Println("Main function is doing other work...")
	}
}
