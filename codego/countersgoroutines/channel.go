package main

import (
	"fmt"
	"sync"
)

func main() {
	n := 20
	evenCh := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range evenCh {
			fmt.Println(num)
		}
	}()

	// send even numbers
	for i := 2; i <= n; i += 2 {
		evenCh <- i
	}
	close(evenCh)

	wg.Wait()
	fmt.Println("Done (channel based)")
}
