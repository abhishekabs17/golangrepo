package main

import (
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

	done := make(chan struct{})

	// Start 1000 goroutines, but let them exit cleanly
	for i := 0; i < 1000; i++ {
		go func(id int) {
			defer fmt.Println("Exiting goroutine:", id)
			select {
			case <-done: // exit when signaled
				return
			case <-time.After(100 * time.Millisecond): // do some "work"
				fmt.Println("Work done by goroutine:", id)
			}
		}(i)
		time.Sleep(10 * time.Millisecond)
	}

	// Give time for all goroutines to finish work
	time.Sleep(2 * time.Second)
	close(done) // signal all goroutines to exit

	fmt.Println("No leaks, all goroutines exit properly")
	select {} // keep main alive for pprof inspection
}
