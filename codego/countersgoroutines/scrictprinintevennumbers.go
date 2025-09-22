package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	n := 20
	counter := 2

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			mu.Lock()
			if counter > n {
				mu.Unlock()
				return
			}
			fmt.Println(counter)
			counter += 2
			mu.Unlock()
		}
	}()

	wg.Wait()
	fmt.Println("Done (strict sequential)")
}
