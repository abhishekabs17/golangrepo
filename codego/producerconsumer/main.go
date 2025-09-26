package main

import (
	"fmt"
	"time"
)

func producer(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
		fmt.Println("Produced:", i)
	}
	close(ch) // important!
}

func consumer(ch chan int) {
	for val := range ch {
		fmt.Println("Consumed:", val)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	ch := make(chan int, 2) // buffered channel
	go producer(ch)
	consumer(ch)
}
