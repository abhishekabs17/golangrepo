Garbage Collection in Go

1.Go uses a concurrent, tricolor mark-and-sweep garbage collector.

2.It pauses briefly to mark live objects, then sweeps unused ones.

3.It runs automatically (you don’t need to free memory manually).

4.You can give hints using runtime.GC() but rarely needed.

takeaway: Go’s GC is optimized for low-latency services and runs concurrently with program execution, minimizing stop-the-world pauses.