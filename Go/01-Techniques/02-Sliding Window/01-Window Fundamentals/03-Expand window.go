/*
Technique: Sliding Window Fundamentals

Pattern: None (Technique Practice)

Operation: Expand Window

Objective:
Understand how a sliding window expands by moving its right boundary
forward.

Key Learning:
The right boundary determines the ending position of the current
window.

When the right boundary moves forward by one position, the element
at the new right position enters the window.

Example:

Before expansion:

[20, 30]
 ↑     ↑
left  right

After right++:

[20, 30, 40]
 ↑        ↑
left     right

Important Concept:
Expanding a window means increasing the right boundary.

The existing elements remain in the window, and the newly reached
element is added to the window.

Operation:
right++

This increases the size of the current window by one element.

Time Complexity: O(1)
Space Complexity: O(1)
*/

package main

import "fmt"

func main() {

	arr := [6]int{10, 20, 30, 40, 50, 60}

	left := 1
	right := 2

	fmt.Println("Before expansion")
	for i := left; i <= right; i++ {
		fmt.Println(arr[i])
	}

	fmt.Println("After expansion")
	right++
	for i := left; i <= right; i++ {
		fmt.Println(arr[i])
	}
}
