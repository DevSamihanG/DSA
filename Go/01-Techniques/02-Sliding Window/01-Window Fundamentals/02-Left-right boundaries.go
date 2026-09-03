/*
Technique: Sliding Window Fundamentals

Pattern: None (Technique Practice)

Operation: Left / Right Boundaries

Objective:
Understand how left and right boundaries define the current window
within an array.

Key Learning:
The left boundary represents where the window starts.
The right boundary represents where the window ends.

Every element between these two boundaries belongs to the window.

If:
    left = 1
    right = 3

Then the window contains:

    arr[1], arr[2], arr[3]

Moving the right boundary forward expands the window.

Moving the left boundary forward shrinks the window from the left.

Important Concept:
The boundaries do not represent only two elements.

They define the entire range between them.

Example:

Index:   0    1    2    3    4    5
Value:  10   20   30   40   50   60
              ↑         ↑
            left       right

left = 1
right = 3

Window:
[20, 30, 40]

After right++:

[20, 30, 40, 50]

After left++:

[30, 40, 50]

Time Complexity: O(1)
Space Complexity: O(1)
*/

/*
Technique: Sliding Window Fundamentals

Pattern: None (Technique Practice)

Operation: Left / Right Boundaries

Objective:
Understand how left and right boundaries define the current window
within an array.

Key Learning:
The left boundary represents where the window starts.
The right boundary represents where the window ends.

Every element between left and right belongs to the window.

Important Concept:
Moving the right boundary forward expands the window.
Moving the left boundary forward shrinks the window from the left.

Time Complexity: O(n)
Space Complexity: O(1)
*/

package main

import "fmt"

func main() {

	arr := [6]int{10, 20, 30, 40, 50, 60}

	left := 1
	right := 3

	// Current window
	fmt.Println(arr[1], arr[2], arr[3])

	// Move right boundary
	right++
	fmt.Println(arr[1], arr[2], arr[3], arr[4])

	// Move left boundary
	left++
	fmt.Println(arr[2], arr[3], arr[4])
}
