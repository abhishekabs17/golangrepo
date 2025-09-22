package main

import (
	"fmt"
)

// Deque is a generic double-ended queue (FIFO and LIFO).
// It uses a slice to store its elements.
type Deque[T any] struct {
	Array []T
}

// PushBack adds an element to the back of the deque.
// This is an efficient O(1) operation on average.
func (d *Deque[T]) PushBack(val T) {
	d.Array = append(d.Array, val)
}

// PushFront adds an element to the front of the deque.
// This is an O(n) operation because all existing elements must be shifted.
func (d *Deque[T]) PushFront(val T) {
	d.Array = append([]T{val}, d.Array...)
}

// PopFront removes and returns the element from the front of the deque.
// This is an O(n) operation because all subsequent elements must be shifted.
// It returns a zero value and false if the deque is empty.
func (d *Deque[T]) PopFront() (T, bool) {
	if d.IsEmpty() {
		var zero T
		return zero, false
	}
	// Get the front element.
	val := d.Array[0]
	// Re-slice the array to remove the first element.
	d.Array = d.Array[1:]
	return val, true
}

// PopBack removes and returns the element from the back of the deque.
// This is an efficient O(1) operation.
// It returns a zero value and false if the deque is empty.
func (d *Deque[T]) PopBack() (T, bool) {
	if d.IsEmpty() {
		var zero T
		return zero, false
	}
	// Get the last element.
	val := d.Array[len(d.Array)-1]
	// Slice off the last element.
	d.Array = d.Array[:len(d.Array)-1]
	return val, true
}

// PeekFront returns the front element without removing it.
// It returns a zero value and false if the deque is empty.
func (d *Deque[T]) PeekFront() (T, bool) {
	if d.IsEmpty() {
		var zero T
		return zero, false
	}
	return d.Array[0], true
}

// PeekBack returns the back element without removing it.
// It returns a zero value and false if the deque is empty.
func (d *Deque[T]) PeekBack() (T, bool) {
	if d.IsEmpty() {
		var zero T
		return zero, false
	}
	return d.Array[len(d.Array)-1], true
}

// IsEmpty checks if the deque is empty.
func (d *Deque[T]) IsEmpty() bool {
	return len(d.Array) == 0
}

// Size returns the number of elements in the deque.
func (d *Deque[T]) Size() int {
	return len(d.Array)
}

func main() {
	// Create a deque of integers
	d := Deque[int]{}

	fmt.Println("Is empty?", d.IsEmpty()) // true

	// Push elements at back
	d.PushBack(10)
	d.PushBack(20)
	d.PushBack(30)
	fmt.Println("After PushBack:", d.Array)

	// Push elements at front
	d.PushFront(5)
	d.PushFront(1)
	fmt.Println("After PushFront:", d.Array)

	// Peek front and back
	front, _ := d.PeekFront()
	back, _ := d.PeekBack()
	fmt.Println("PeekFront:", front) // 1
	fmt.Println("PeekBack:", back)   // 30

	// Pop from front
	val, ok := d.PopFront()
	if ok {
		fmt.Println("PopFront:", val)
	}
	fmt.Println("Deque after PopFront:", d.Array)

	// Pop from back
	val, ok = d.PopBack()
	if ok {
		fmt.Println("PopBack:", val)
	}
	fmt.Println("Deque after PopBack:", d.Array)

	// Size of deque
	fmt.Println("Size:", d.Size())
}
