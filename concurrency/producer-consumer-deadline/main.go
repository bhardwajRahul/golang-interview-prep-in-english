package main

import (
	"context"
	"fmt"
	"sort"
	"sync"
	"time"
)

func producer(mu *sync.Mutex, ctx context.Context, ch chan<- int, wg *sync.WaitGroup, index int) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		// Symulujemy pracę producenta.
		select {
		case <-ctx.Done():
			fmt.Printf("Producer %d stopped: %v\n", index, ctx.Err())
			return
		case <-time.After(20 * time.Millisecond):
		}

		value := index*10 + i

		select {
		case <-ctx.Done():
			fmt.Printf("Producer %d stopped: %v\n", index, ctx.Err())
			return
		case ch <- value:
			fmt.Println("Sent:", value)
		}
	}
}

func consumer(ch <-chan int, wg *sync.WaitGroup, index int) {
	defer wg.Done()
	var values []int

	for v := range ch {
		values = append(values, v)
	}

	sort.Ints(values)

	for _, v := range values {
		fmt.Printf("Consumer %d received: %d\n", index, v)
	}
}
func main() {
	ch := make(chan int, 10)
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	wp := sync.WaitGroup{}
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go producer(&mu, ctx, ch, &wg, i)
	}
	for i := 0; i < 10; i++ {
		wp.Add(1)

		go consumer(ch, &wp, i)
	}
	wg.Wait()
	close(ch)
	wp.Wait()

	println("Done")
}
