/*
Technique: Sliding Window

Pattern: None (Technique Practice)

Operation: Running Count

Objective:
Understand how to maintain a count of elements inside a sliding window
that satisfy a specific condition.

Key Learning:
A running count stores how many elements in the current window satisfy
a given condition.

First, calculate the count for the initial window.

When the window expands:
A new element enters from the right. Check whether the new element
satisfies the condition.

If it does:
    count++

When the window shrinks:
The element at the left boundary leaves the window. Check whether that
element satisfies the condition before moving the left boundary.

If it does:
    count--

Example:

Condition:
Count even numbers.

Initial window:

[10, 20, 30]

All three elements are even.

count = 3

After expansion:

[10, 20, 30, 40]

40 enters the window.
40 is even.

count = 4

After shrinking:

[20, 30, 40]

10 leaves the window.
10 is even.

count = 3

Important Concept:
Only elements that satisfy the required condition affect the count.

When an element enters:

If it satisfies the condition:
    count++

When an element leaves:

If it satisfies the condition:
    count--

Time Complexity: O(n)
Space Complexity: O(1)
*/

package main

import (
	"fmt"
)

func main() {

	arr := [6]int{10, 20, 30, 40, 50, 60}

	left := 0
	right := 2
	count := 0

	// Calculate count for initial window
	for i := left; i <= right; i++ {
		if arr[i]%2 == 0 {
			count++
		}
	}

	fmt.Println("Initial window count:", count)

	// Expand window: new right element enters
	right++
	if arr[right]%2 == 0 {
		count++
	}

	fmt.Println("After window expansion:", count)

	// Shrink window: old left element leaves
	if arr[left]%2 == 0 {
		count--
	}
	left++

	fmt.Println("After window shrinking:", count)
}
