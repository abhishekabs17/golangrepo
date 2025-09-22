package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	n := 20
	counter := 2

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for counter <= n {
			mu.Lock()
			fmt.Println(counter)
			counter += 2
			cond.Signal() // wake another goroutine if waiting
			mu.Unlock()
		}
	}()

	wg.Wait()
	fmt.Println("Done (Cond based)")
}
