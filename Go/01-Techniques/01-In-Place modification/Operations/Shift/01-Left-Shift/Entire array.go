package main

import (
	"fmt"
)

func main() {

	arr := [5]int{5, 8, 7, 3, 2}

	for i := 0; i < len(arr)-1; i++ {
		arr[i] = arr[i+1]
	}

	fmt.Println(arr)
}
