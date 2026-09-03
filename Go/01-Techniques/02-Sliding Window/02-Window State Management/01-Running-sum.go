/*
Technique: Sliding Window

Pattern: None (Technique Practice)

Operation: Running Sum

Objective:
Understand how to maintain the sum of elements currently inside a
sliding window without recalculating the entire sum every time the
window changes.

Key Learning:
A running sum stores the current sum of all elements inside the window.

First, calculate the sum of the initial window.

When the window expands:
A new element enters from the right, so add that element to the
existing sum.

When the window shrinks:
The element at the left boundary leaves the window, so subtract that
element from the existing sum.

Example:

Initial window:

[10, 20, 30]

Sum:
10 + 20 + 30 = 60

After expansion:

[10, 20, 30, 40]

New element entering:
40

Updated sum:
60 + 40 = 100

After shrinking:

[20, 30, 40]

Element leaving:
10

Updated sum:
100 - 10 = 90

Important Concept:
Do not recalculate the sum of the entire window after every change.

Instead:

Element enters window → Add it to sum
Element leaves window → Subtract it from sum

This allows the window state to be updated efficiently as the
boundaries move.

Time Complexity: O(n)
Space Complexity: O(1)
*/

package main

import "fmt"

func main() {

	arr := [6]int{10, 20, 30, 40, 50, 60}

	left := 0
	right := 2
	sum := 0

	// Calculate initial window sum
	for i := left; i <= right; i++ {
		sum += arr[i]
	}

	fmt.Println("Initial sum:", sum)

	// Expand window: new right element enters
	right++
	sum += arr[right]

	fmt.Println("After expansion:", sum)

	// Shrink window: old left element leaves
	sum -= arr[left]
	left++

	fmt.Println("After shrinking:", sum)
}
