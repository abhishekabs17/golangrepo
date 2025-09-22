package main

import (
	"fmt"
	"sync"
)

func main() {
	n := 20
	oddCh := make(chan bool)
	evenCh := make(chan bool)
	var wg sync.WaitGroup

	wg.Add(2)

	// Even printer
	go func() {
		defer wg.Done()
		for i := 2; i <= n; i += 2 {
			<-evenCh
			fmt.Println("Even:", i)
			if i+1 <= n {
				oddCh <- true
			} else {
				close(oddCh) // stop odd goroutine gracefully
			}
		}
	}()

	// Odd printer
	go func() {
		defer wg.Done()
		for i := 1; i <= n; i += 2 {
			val, ok := <-oddCh
			if !ok {
				return
			}
			_ = val
			fmt.Println("Odd:", i)
			evenCh <- true
		}
	}()

	// Start with odd
	oddCh <- true

	wg.Wait()
	fmt.Println("Done (odd/even alternate, no deadlock)")
}
