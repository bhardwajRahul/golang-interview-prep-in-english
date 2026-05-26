package main

import (
	"fmt"
	"sync"
)

func fanOut(in <-chan int, n int) []<-chan int {
	outs := make([]<-chan int, n)
	for i := 0; i < n; i++ {
		ch := make(chan int)
		outs[i] = ch
		go func() {
			defer close(ch)
			for v := range in {
				ch <- v * v // Square it
			}
		}()
	}
	return outs
}

func fanIn(chs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(chs))
	for _, ch := range chs {
		go func(c <-chan int) {
			defer wg.Done()
			for v := range c {
				out <- v
			}
		}(ch)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func pipeline(input <-chan int) <-chan int {
	workers := fanOut(input, 3) // 3 workers
	return fanIn(workers...)    // Merge results
}

func generateInput() <-chan int {
	in := make(chan int)
	go func() {
		defer close(in)
		for i := 1; i <= 5; i++ {
			in <- i
		}
	}()
	return in
}

func main() {
	input := generateInput()
	result := pipeline(input)
	for v := range result {
		fmt.Println(v) // 1, 4, 9, 16, 25 (order varies)
	}
}
