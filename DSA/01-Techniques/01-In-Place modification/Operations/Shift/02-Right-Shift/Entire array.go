package main

import (
	"fmt"
)

func main() {

	arr := [5]int{5, 8, 3, 7, 2}

	for i := len(arr) - 1; i > 0; i-- {
		arr[i] = arr[i-1]

	}
	fmt.Println(arr)
}
