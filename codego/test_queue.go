package main

import "fmt"

type queue[T any] struct {
	array []T
}

func (q *queue[T]) dequeue() (T, bool) {
	var zero T
	if len(q.array) == 0 {
		return zero, false
	}
	val := q.array[0]
	q.array = q.array[1:]
	return val, true
}

func (q *queue[T]) enqueue(val T) {
	q.array = append(q.array, val)
}

func (q *queue[T]) peek() (T, bool) {
	if len(q.array) != 0 {
		return q.array[0], true
	}
	var zero T
	return zero, false
}

func (q *queue[T]) size() int {
	a := len(q.array)
	return a
}

func (q *queue[T]) isempty() bool {
	return len(q.array) == 0
}

// Iterate over queue
func (q *queue[T]) Iterate(f func(T)) {
	for _, val := range q.array {
		f(val)
	}
}

func main() {
	var q queue[int]
	q.enqueue(10)
	q.enqueue(20)
	q.enqueue(30)

	fmt.Println(q.size())
	fmt.Println(q.peek())

	val, _ := q.dequeue()
	fmt.Println(val)
	fmt.Println(q.size())

	// Iterate over queue

	// for _, val := range q.array {
	// 	fmt.Println(val)
	// }

	q.Iterate(func(val int) {
		fmt.Println(val)
	})
}
