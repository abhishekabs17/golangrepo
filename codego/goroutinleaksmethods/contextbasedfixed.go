package main

import (
	"context"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	// Start pprof server
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	// Create a context that will cancel after 2 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Launch goroutines that respect context cancellation
	for i := 0; i < 1000; i++ {
		go func(id int) {
			select {
			case <-ctx.Done(): // canceled when timeout expires
				fmt.Println("Goroutine", id, "exiting due to cancel:", ctx.Err())
				return
			case <-time.After(100 * time.Millisecond): // simulate some work
				fmt.Println("Work done by goroutine:", id)
			}
		}(i)
		time.Sleep(10 * time.Millisecond)
	}

	fmt.Println("Launched goroutines with context cancellation")
	select {} // keep main alive for pprof
}
