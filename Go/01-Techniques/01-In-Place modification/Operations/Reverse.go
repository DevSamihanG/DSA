/*
Technique: Reverse (In-Place Modification)

Idea:
- Reversing cannot be done using overwrite because values would be lost.
- We exchange (swap) the elements from both ends of the array.
- Two pointers are used:
  - left starts from the beginning.
  - right starts from the end.

Algorithm:
1. Initialize:
   left = 0
   right = last index

2. While left < right:
   - Swap arr[left] and arr[right].
   - Move left one step right (left++).
   - Move right one step left (right--).

3. Stop when left >= right.
   At this point, every element has reached its reversed position.

Time Complexity : O(n)
Space Complexity: O(1)

Key Learning:
- This is an in-place algorithm.
- Reverse is built using repeated swaps.
- This is the first example of the Two Pointer technique where
  two pointers move towards each other independently.
*/

package main

import (
	"fmt"
)

func main() {

	arr := [6]int{10, 20, 30, 40, 50, 60}

	left := 0
	right := len(arr) - 1

	for left < right {
		temp := arr[left]
		arr[left] = arr[right]
		arr[right] = temp
		left++
		right--
	}
	fmt.Println(arr)
}
