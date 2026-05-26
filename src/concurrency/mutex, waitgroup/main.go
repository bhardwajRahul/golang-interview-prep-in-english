package main

import (
	"fmt"
	"sync"
)

func main() {
	mu := sync.Mutex{}
	counter := 0
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1) // Increment the WaitGroup counter for each goroutine
		go func() {
			defer wg.Done()   // Decrement the WaitGroup counter when the goroutine completes
			mu.Lock()         // Lock the mutex before accessing the counter
			defer mu.Unlock() // Ensure the mutex is unlocked after the function returns
			counter++
		}()
	}
	wg.Wait() // Wait for all goroutines to finish
	// Wait for a moment to let goroutines finish (in a real application, use sync.WaitGroup)
	fmt.Println("Counter value:", counter) // This may not reflect the correct count due to
}
