Escape Analysis in Go

Definition:
Escape analysis is a compiler optimization technique in Go that determines whether a variable can be safely allocated on the stack or must "escape" to the heap.

Stack allocation → faster, automatically cleaned up when function returns.

Heap allocation → slower (needs garbage collection), but survives beyond function scope.

👉 Rule of thumb:

If a variable’s lifetime ends within the function → allocated on stack.

If a variable/reference is returned or used outside → escapes to heap.


Memory Allocation (Stack vs Heap)

1.Stack → used for local variables inside a function (fast allocation, automatically freed).

2.Heap → used when variables escape function scope (e.g., returned, referenced elsewhere).

3.Escape analysis decides whether allocation happens on stack or heap.