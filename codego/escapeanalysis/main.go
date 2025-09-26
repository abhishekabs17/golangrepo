package main

import "fmt"

func noEscape() int {
	x := 10
	return x // stays on stack
}

func escape() *int {
	y := 20
	return &y // escapes to heap
}

func main() {
	a := noEscape()
	b := escape()
	fmt.Println(a, *b)
}
