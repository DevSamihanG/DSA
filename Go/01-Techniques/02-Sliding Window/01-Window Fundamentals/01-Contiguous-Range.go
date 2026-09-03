/*
Technique: Sliding Window Fundamentals

Pattern: None (Technique Practice)

Operation: Contiguous Range

Objective:
Understand how a contiguous range is represented using two boundaries
in an array.

Key Learning:
A contiguous range contains every element between a starting index
and an ending index.

If left = 1 and right = 4, the range contains indices:

1 → 2 → 3 → 4

No elements inside the range can be skipped.

Important Concept:
The two boundaries define the range:

left  = starting position
right = ending position

Everything between left and right belongs to the range.

A contiguous range:
[20, 30, 40, 50]

A non-contiguous selection:
[20, 40, 50]

The second selection skips index 2, so it is not contiguous.

Time Complexity: O(n)
Space Complexity: O(1)
*/

package main

import "fmt"

func main() {

	arr := [6]int{10, 20, 30, 40, 50, 60}

	left := 1
	right := 4

	fmt.Println("Contiguous range:")

	for i := left; i <= right; i++ {
		fmt.Println(arr[i])
	}
}
