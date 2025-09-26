sync/atomic → Used for very simple counters or flags, where operation is small (e.g., increment, compare-and-swap).

Faster than mutex (lock-free).

sync.Mutex → Use when you need to protect larger critical sections or multiple operations.


What this demonstrates

1.sync.Once → initOnce only runs once.

2.sync.Mutex → incrementWithMutex safely increments counter.

3.sync.RWMutex → multiple readers allowed, writer blocks them.

4.sync.WaitGroup → waits for goroutines to finish.

5.sync.Cond → goroutines wait until condition is signaled/broadcast.