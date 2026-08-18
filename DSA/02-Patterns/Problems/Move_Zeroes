/*
Move Zeroes

Idea:
Use two pointers:
- Read Pointer  -> Traverses every element.
- Write Pointer -> Tracks the next position for a non-zero element.

Algorithm:
1. Read pointer scans the array from left to right.
2. If a non-zero element is found, overwrite the Write Pointer position and move Write forward.
3. After all non-zero elements are placed at the front, fill the remaining positions with zeros.

Reason:
The Write Pointer always maintains the boundary of the correct prefix.
After the first pass, all non-zero elements are in their correct relative order.
The remaining positions are no longer needed and can safely be replaced with zeros.
*/

package main

import (
	"fmt"
)

func main() {
	// 0  1  2  3   4
	arr := [5]int{0, 1, 0, 3, 12}

	write := 0

	for read := 0; read < len(arr); read++ {
		if arr[read] != 0 {
			arr[write] = arr[read]
			write++
		}

	}
	for i := write; i < len(arr); i++ {
		arr[i] = 0
	}
	fmt.Println(arr)
}
