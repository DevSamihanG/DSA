package main

import (
	"fmt"
)

func main() {
	// 0  1  2  3  4  5  6  7
	arr := [8]int{0, 1, 2, 2, 3, 0, 4, 2}

	val := 2

	write := 0

	for read := 0; read < len(arr); read++ {
		if arr[read] != val {
			arr[write] = arr[read]
			write++
		}
	}
	fmt.Println(arr)
	fmt.Println(write)
}
