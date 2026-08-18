package main

import (
	"fmt"
)

func main() {

	arr := [6]int{2, 4, 7, 11, 15, 20}
	target := 26

	right := 0
	left := len(arr) - 1

	for left < right {

		sum := arr[left] + arr[right]

		if sum < target {
			left++
		} else if sum > target {
			right--
		} else {
			fmt.Println(arr[left], arr[right])
			return
		}
	}

}
