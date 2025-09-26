package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, 世界")
	arr := []int{2, 3, 4, 5, 6}
	arr1 := arr[2:]
	arr1[2] = 10
	fmt.Println("arr", arr)
	arr2 := arr1[1:]
	arr2[0] = 17
	fmt.Println("arr after change", arr)
}
