package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	leakch := make(chan int)
	go func() {
		http.ListenAndServe("localhost:8080", nil)
	}()

	for i := 0; i < 1000; i++ {
		go func(id int) {
			fmt.Println("goroutine id:", id)
			<-leakch
		}(i)
		time.Sleep(5 * time.Millisecond)
	}

	fmt.Println("Leaked 1000 goroutines")
	select {}
}
