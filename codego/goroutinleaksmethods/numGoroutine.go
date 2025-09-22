package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Printf("no of cpu %d\n", runtime.NumCPU())
		fmt.Printf("no of cgocall %d\n", runtime.NumCgoCall())
		//ch <- 1
		time.Sleep(2 * time.Second)
	}()
	fmt.Printf("Active goroutines1: %d\n", runtime.NumGoroutine())
	wg.Wait()
	fmt.Printf("Active goroutines: %d\n", runtime.NumGoroutine())
	fmt.Println("finished")
}
