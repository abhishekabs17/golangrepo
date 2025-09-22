package main

import (
	"fmt"
	"time"
)

func worker(id int, tasks <-chan int, result chan<- int) {
	for task := range tasks {
		fmt.Printf("worker number %d,processing task%d\n", id, task)
		time.Sleep(time.Second)
		result <- task * task
	}
}

func main() {
	numWorkers := 3
	numTasks := 5

	tasks := make(chan int, numTasks)
	results := make(chan int, numTasks)

	// Start the workers
	for i := 1; i <= numWorkers; i++ {
		go worker(i, tasks, results)
	}

	for j := 1; j <= numTasks; j++ {
		tasks <- j
	}
	close(tasks)

	for k := 1; k <= numTasks; k++ {
		result := <-results
		fmt.Println(result)
	}

	close(results)
}
