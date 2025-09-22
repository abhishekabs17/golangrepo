package main

import (
	"fmt"
)

// Generic Stack
type Stack[T any] struct {
	Array      []T
	StackSize  int
	TopPointer int
}

// add element at top of stack
func (s *Stack[T]) Push(val T) {
	if s.TopPointer == s.StackSize {
		s.Array = append(s.Array, val)
		s.StackSize++
		s.TopPointer++
	} else {
		s.Array[s.TopPointer] = val
		s.TopPointer++
	}
}

// removes top element from stack
func (s *Stack[T]) Pop() {
	if s.TopPointer != 0 {
		s.TopPointer--
	}
}

// IsEmpty checks if the stack is empty.
func (s *Stack[T]) IsEmpty() bool {
	return s.TopPointer == 0
}

// return top element and true if stack is not empty
// else return default value and false
func (s *Stack[T]) Top() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}
	return s.Array[s.TopPointer-1], true
}

// return no of element in stack
func (s *Stack[T]) Size() int {
	return s.TopPointer
}

// return iterator to iterate over stack in top down fashion
func (s *Stack[T]) Iterator() <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		for i := s.TopPointer - 1; i >= 0; i-- {
			ch <- s.Array[i]
		}
	}()
	return ch
}

func main() {
	// Example with integers
	var st Stack[int]

	// Push values
	st.Push(10)
	st.Push(20)
	st.Push(30)

	fmt.Println("Stack Size:", st.Size())

	// Top element
	if top, ok := st.Top(); ok {
		fmt.Println("Top Element:", top)
	}

	// Pop one element
	st.Pop()
	if top, ok := st.Top(); ok {
		fmt.Println("After Pop, Top Element:", top)
	}

	// Iterate over stack
	fmt.Println("Iterating stack (top to bottom):")
	for val := range st.Iterator() {
		fmt.Println(val)
	}

	// Example with strings
	var stStr Stack[string]
	stStr.Push("Go")
	stStr.Push("is")
	stStr.Push("awesome")

	fmt.Println("\nString Stack Iteration:")
	for val := range stStr.Iterator() {
		fmt.Println(val)
	}
}
