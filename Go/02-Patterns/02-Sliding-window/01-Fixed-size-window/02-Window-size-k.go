/*
Technique: Fixed-Size Sliding Window

Pattern: Sliding Window

Operation: Window Size K

Objective:
Understand how to maintain and process a contiguous window
of exactly size k while moving through an array.

Key Learning:
Start with the left boundary at the beginning of the array.

Move the right boundary forward one position at a time.

As right moves, the window expands.

The current window size is calculated as:

```
right - left + 1
```

If the window size is smaller than k, continue expanding
the window.

When the window size becomes exactly k, the window is complete
and can be processed.

After processing the window, move the left boundary forward.

The next iteration moves right forward automatically, causing
the window to slide forward while maintaining size k.

Example:

Array:

[10, 20, 30, 40, 50, 60]

k = 3

Window progression:

[10]                Size < k → Continue expanding

[10, 20]            Size < k → Continue expanding

[10, 20, 30]        Size == k → Process

Move left forward.

[20, 30, 40]        Size == k → Process

Move left forward.

[30, 40, 50]        Size == k → Process

Move left forward.

[40, 50, 60]        Size == k → Process

Important Concept:
The window does not start at size k.

It grows as the right boundary moves forward.

Once the window reaches size k, left moves forward after
each processed window while right continues moving forward.

This keeps every processed window at exactly size k.

Window Movement:

left = 0, right = 0
[10]

right moves forward
[10, 20]

right moves forward
[10, 20, 30] → Process

left moves forward
right moves forward

[20, 30, 40] → Process

Time Complexity: O(n × k) for this demonstration,
because every element inside each window is printed.

Space Complexity: O(1)
*/

package main

import "fmt"

func main() {

	arr := [6]int{10, 20, 30, 40, 50, 60}

	k := 3
	left := 0

	for right := 0; right < len(arr); right++ {

		windowSize := right - left + 1

		if windowSize == k {
			for i := left; i <= right; i++ {
				fmt.Print(arr[i])
			}
			left++
		}
	}
}
