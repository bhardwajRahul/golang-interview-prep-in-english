package main

func main() {
	ch := make(chan int, 5) // Create a buffered channel with capacity of 2
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
	for v := range ch {
		println(v)
	}
	println("Done")
}
