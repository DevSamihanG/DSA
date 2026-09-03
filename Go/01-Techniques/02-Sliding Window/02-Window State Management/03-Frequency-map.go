/*
Technique: Sliding Window

Pattern: None (Technique Practice)

Operation: Frequency Map

Objective:
Understand how to maintain the frequency of elements currently inside
a sliding window.

Key Learning:
A frequency map stores how many times each element appears inside the
current window.

The map follows the structure:

element → number of occurrences

First, build the frequency map for the initial window.

When the window expands:
A new element enters from the right, so increase its frequency.

When the window shrinks:
The element at the left boundary leaves the window, so decrease its
frequency.

If the frequency of that element becomes 0, remove it from the map.

Example:

Initial window:

[10, 20, 30]

Frequency Map:

10 → 1
20 → 1
30 → 1

After expansion:

[10, 20, 30, 40]

40 enters the window.

Frequency Map:

10 → 1
20 → 1
30 → 1
40 → 1

After shrinking:

[20, 30, 40]

10 leaves the window.

First:

frequency[10]--

10 → 0

Since its frequency is now 0:

delete(frequency, 10)

Final Frequency Map:

20 → 1
30 → 1
40 → 1

Important Concept:

When an element enters the window:

frequency[element]++

When an element leaves the window:

frequency[element]--

If its frequency becomes 0:

delete(frequency, element)

This keeps the frequency map synchronized with the elements currently
present inside the window.

Time Complexity: O(n)
Space Complexity: O(n)
*/

package main

import (
	"fmt"
)

func main() {

	arr := [6]int{10, 20, 30, 40, 50, 60}

	left := 0
	right := 2

	frequency := make(map[int]int)

	// Build frequency map for initial window
	for i := left; i <= right; i++ {
		frequency[arr[i]]++
	}

	fmt.Println("Initial map frequency:", frequency)

	// Expand window: new right element enters
	right++
	frequency[arr[right]]++

	fmt.Println("Map frequency after expansion:", frequency)

	// Shrink window: old left element leaves
	frequency[arr[left]]--

	if frequency[arr[left]] == 0 {
		delete(frequency, arr[left])
	}

	left++

	fmt.Println("Map frequency after shrinking:", frequency)
}
