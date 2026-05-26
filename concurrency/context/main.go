package main

import (
	"context"
)

func main() {
	// context propagates cancellation signals and values across API boundaries and goroutines.
	ctx, cancel := context.WithCancel(context.WithValue(context.Background(), "key", 100))
	go func() {
		cancel() // Cancel the context after some work is done
	}()
	select {
	case <-ctx.Done():
		println("Context cancelled")
		value := ctx.Value("key")
		println("Value from context:", value.(int))
	}
}
