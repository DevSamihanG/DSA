/*
Technique: Running Sum with Fixed-Size Sliding Window

Pattern: Sliding Window

Operation: Processing a Window

Objective:
Understand how to efficiently process every contiguous window
of size k by maintaining its sum.

Problem:
Print the sum of every contiguous window of size k.

Array:

[10, 20, 30, 40, 50, 60]

k = 3

Expected Windows and Sums:

[10, 20, 30] → 60
[20, 30, 40] → 90
[30, 40, 50] → 120
[40, 50, 60] → 150

Key Learning:
Instead of calculating the sum of every window from scratch,
maintain a running sum as the window moves.

When the right boundary moves forward, a new element enters
the window.

Add that element to the running sum:

```
sum += arr[right]
```

The current window size is:

```
right - left + 1
```

When the window size becomes exactly k, the running sum
represents the sum of the complete current window.

Process the window:

```
fmt.Println(sum)
```

Before moving to the next window, the element at the left
boundary must leave the current window.

Remove it from the running sum:

```
sum -= arr[left]
```

Then move the left boundary forward:

```
left++
```

Window Movement:

Initial:

left = 0
right = 0

[10]

sum = 10

Expand:

right = 1

[10, 20]

sum = 30

Expand:

right = 2

[10, 20, 30]

sum = 60

windowSize == k

Process:

60

Prepare for the next window:

Remove arr[left] = 10

sum = 60 - 10 = 50

Move left:

left++

Next Iteration:

right = 3

40 enters the window.

sum = 50 + 40 = 90

Window:

[20, 30, 40]

windowSize == k

Process:

90

Important Concept:
The sum is not recalculated from the beginning for every
new window.

The window maintains its previous state.

When a new element enters:

```
Add it to the sum.
```

When an old element leaves:

```
Remove it from the sum.
```

This allows the window to move efficiently through the array.

General Flow:

right moves forward
↓
New element enters
↓
Add arr[right] to sum
↓
Calculate window size
↓
windowSize == k?
│
├── No
│     ↓
│   Continue expanding
│
└── Yes
↓
Process sum
↓
Remove arr[left] from sum
↓
left++
↓
Next iteration moves right forward

Time Complexity: O(n)

Each element is added to the sum once and removed from
the sum at most once.

Space Complexity: O(1)

Only a fixed number of variables are used.
*/

package main

import (
	"fmt"
)

func main() {

	arr := [6]int{10, 20, 30, 40, 50, 60}

	k := 3
	left := 0
	sum := 0

	for right := 0; right < len(arr); right++ {

		// Add the new element entering from the right.
		sum += arr[right]

		// Calculate the current window size.
		windowSize := right - left + 1

		// Process the window once it reaches size k.
		if windowSize == k {

			fmt.Println(sum)

			// Remove the element leaving from the left.
			sum -= arr[left]

			// Move the left boundary forward.
			left++
		}
	}
}
