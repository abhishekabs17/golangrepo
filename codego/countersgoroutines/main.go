package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	n := 20 // print even numbers up to 20

	// shared counter
	counter := 2

	// goroutine function
	printEven := func(id int) {
		defer wg.Done()
		for {
			mu.Lock()
			if counter > n {
				mu.Unlock()
				return
			}
			// print and increase
			fmt.Printf("Goroutine %d -> %d\n", id, counter)
			counter += 2
			mu.Unlock()
		}
	}

	// let's run 3 goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go printEven(i)
	}

	wg.Wait()
	fmt.Println("Finished printing even numbers")
}
