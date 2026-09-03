/*
Technique: Pointer Technique

Pattern: None (Technique Practice)

Operation: Left / Right Pointers

Objective:
Learn how two pointers can start at opposite ends of an array and
move toward each other.

Key Learning:
The left pointer starts at the beginning of the array.
The right pointer starts at the end of the array.

Both pointers represent positions in the array.

The pointers move inward based on the condition of the problem.

Important Concept:
Left and right pointers allow us to work with elements from both
ends of an array without creating another array.

The pointers continue moving until they meet or cross.

This technique is commonly used for:
- Comparing elements from both ends
- Searching for pairs
- Reversing arrays
- Partitioning
- Container / two-pointer problems

Time Complexity: O(n)
Space Complexity: O(1)
*/

package main

import "fmt"

func main() {

	arr := [5]int{10, 20, 30, 40, 50}

	left := 0
	right := len(arr) - 1

	for left < right {

		fmt.Println("Left:", arr[left], "Right:", arr[right])

		left++
		right--
	}
}
