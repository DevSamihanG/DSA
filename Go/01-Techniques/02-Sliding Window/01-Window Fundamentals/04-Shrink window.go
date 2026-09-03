/*
Technique: Sliding Window Fundamentals

Pattern: None (Technique Practice)

Operation: Shrink Window

Objective:
Understand how a sliding window shrinks by moving its left boundary
forward.

Key Learning:
The left boundary determines the starting position of the current
window.

When the left boundary moves forward by one position, the element
at the old left position leaves the window.

Example:

Before shrinking:

[20, 30, 40]
 ↑          ↑
left       right

After left++:

[30, 40]
 ↑       ↑
left    right

Important Concept:
Shrinking a window means increasing the left boundary.

The remaining elements stay in the window, while the element at the
old left position is removed from the window.

Operation:
left++

This decreases the size of the current window by one element.

Time Complexity: O(n)
Space Complexity: O(1)
*/

package main

import "fmt"

func main() {

	arr := [6]int{10, 20, 30, 40, 50, 60}

	left := 1
	right := 3

	fmt.Println("Before shrinking")
	for i := left; i <= right; i++ {
		fmt.Println(arr[i])
	}

	fmt.Println("After shrinking")
	left++
	for i := left; i <= right; i++ {
		fmt.Println(arr[i])
	}
}
