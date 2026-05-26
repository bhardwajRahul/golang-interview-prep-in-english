package main

import (
	"fmt"
	"sync"
)

func fanIn(chs ...<-chan string) <-chan string {
	out := make(chan string)
	var wg sync.WaitGroup
	wg.Add(len(chs))

	for _, ch := range chs {
		go func(c <-chan string) {
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

func generateData(id int) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for i := 0; i < 3; i++ {
			ch <- fmt.Sprintf("Worker %d: %d", id, i)
		}
	}()
	return ch
}

func main() {
	ch1, ch2, ch3 := generateData(1), generateData(2), generateData(3)
	result := fanIn(ch1, ch2, ch3)
	for v := range result {
		fmt.Println(v)
	}
}
