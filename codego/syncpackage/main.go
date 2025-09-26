package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	counter int
	mu      sync.Mutex   // for exclusive lock
	rwmu    sync.RWMutex // for read-write lock
	once    sync.Once    // for one-time init
	cond    = sync.NewCond(&sync.Mutex{})
)

func initOnce() {
	fmt.Println("Initialized only once")
}

func incrementWithMutex(wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	defer mu.Unlock()
	counter++
	fmt.Println("Counter incremented:", counter)
}

func readWithRWMutex(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	rwmu.RLock()
	defer rwmu.RUnlock()
	fmt.Printf("Reader %d sees counter = %d\n", id, counter)
}

func writeWithRWMutex(wg *sync.WaitGroup) {
	defer wg.Done()
	rwmu.Lock()
	defer rwmu.Unlock()
	counter += 10
	fmt.Println("Writer updated counter to:", counter)
}

func condWaiter(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	cond.L.Lock()
	fmt.Printf("Waiter %d waiting...\n", id)
	cond.Wait() // wait until someone signals
	fmt.Printf("Waiter %d resumed\n", id)
	cond.L.Unlock()
}

func main() {
	var wg sync.WaitGroup

	// sync.Once example
	once.Do(initOnce)
	once.Do(initOnce) // won't run again

	// sync.Mutex example
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go incrementWithMutex(&wg)
	}
	wg.Wait()

	// sync.RWMutex example
	wg.Add(3)
	go readWithRWMutex(1, &wg)
	go readWithRWMutex(2, &wg)
	go writeWithRWMutex(&wg)
	wg.Wait()

	// sync.Cond example
	wg.Add(3)
	go condWaiter(1, &wg)
	go condWaiter(2, &wg)
	go condWaiter(3, &wg)

	time.Sleep(1 * time.Second) // give waiters time to block

	cond.L.Lock()
	fmt.Println("Broadcasting signal to all waiters...")
	cond.Broadcast() // wake up all waiters
	cond.L.Unlock()

	wg.Wait()

	fmt.Println("Final Counter:", counter)
}
