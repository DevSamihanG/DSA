/*
Remove Duplicates from Sorted Array

Idea:
Use two pointers:
- Read Pointer  -> Traverses every element.
- Write Pointer -> Builds the array with unique elements.

Algorithm:
1. The first element is always unique.
2. Read pointer scans the array from left to right.
3. If the current element is different from the previous one,
   move the write pointer and overwrite its position with the new unique value.
4. Return write + 1, since write stores the last unique index.

Reason:
The array is sorted, so duplicate elements are always adjacent.
Comparing the current element with the previous one is enough to detect duplicates.
*/

package main

import (
	"fmt"
)

func main() {
	//0  1  2  3  4  5  6
	arr := [7]int{1, 1, 2, 2, 3, 4, 4}

	write := 0

	for read := 1; read < len(arr); read++ {

		if arr[read] != arr[read-1] {
			write++
			arr[write] = arr[read]
		}
	}
	fmt.Println("Modified Array:", arr)
	fmt.Println("Unique Elements:", write+1)
}
