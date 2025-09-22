package main

import (
	"fmt"
	"sync"
)

func main() {
	n := 20
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)

	go func() {
		defer wg.Done()
		for val := range ch {
			fmt.Println(val)
		}
	}()

	for i := 2; i <= n; i += 2 {
		ch <- i
	}

	close(ch)
	wg.Wait()

	fmt.Println("completed")
}
