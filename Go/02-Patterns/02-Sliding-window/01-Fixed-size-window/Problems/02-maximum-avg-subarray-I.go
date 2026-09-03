/*
LeetCode 643: Maximum Average Subarray I

Pattern:
Fixed-Size Sliding Window

Problem:
Given an integer array nums and an integer k, find the contiguous
subarray of length exactly k that has the maximum average value.

Idea:
Use a fixed-size sliding window.

* right expands the window.
* left marks the start of the window.
* sum maintains the sum of the current window.
* avg stores the average of the current complete window.
* maxAvg stores the maximum average found so far.

Algorithm:

1. Start left at index 0.
2. Move right through the array.
3. Add nums[right] to sum as each element enters the window.
4. Calculate the current window size.
5. When windowSize == k:

   * Calculate the average of the current window.
   * If this is the first complete window, initialize maxAvg.
   * Otherwise, compare avg with maxAvg.
   * Update maxAvg if the current average is larger.
   * Remove nums[left] from sum because it is leaving the window.
   * Move left forward to slide the window.
6. Return maxAvg.

Why do we calculate avg only when windowSize == k?

Before the window reaches size k, it is incomplete.

Example with k = 4:

[1]
[1, 12]
[1, 12, -5]
[1, 12, -5, -6]  <- First complete window

We only calculate the average after the window contains exactly
k elements because only then is it a valid window for the problem.

Why convert sum and k to float64?

sum and k are integers.

If we do:


sum / k


Go performs integer division.

Example:


51 / 4 = 12


The decimal part is lost.

But the problem requires an average, which can contain decimals.

Therefore:


float64(sum) / float64(k)


Example:


float64(51) / float64(4) = 12.75


Why do we use:


var maxAvg float64


instead of:


maxAvg := 0.0


The array can contain negative numbers.

Example:

nums = [-5, -3]
k = 2

The only window average is:

(-5 + -3) / 2 = -4.0

If:


maxAvg := 0.0


Then:


-4.0 > 0.0


is false.

maxAvg incorrectly remains 0.0 even though 0.0 is not the
average of any valid window.

Therefore, we initialize maxAvg using the first actual valid window.

Why do we check:


if left == 0


When the first complete window is found:


left = 0
windowSize = k


Therefore, left == 0 identifies the first complete window.

For the first window:


maxAvg = avg


For every later window:


if avg > maxAvg
    maxAvg = avg


This ensures maxAvg starts with a real answer, whether the values
are positive or negative.

Reason:
Instead of calculating the sum of every window from scratch, we
maintain the sum incrementally.

When a new element enters:


sum += nums[right]


When the window is complete:


Process the current window.


When an old element leaves:


sum -= nums[left]


Then:


left++


This allows the window to slide forward efficiently.

Fixed-Size Sliding Window Framework:

for right := 0; right < len(nums); right++ {


// Add element entering the window

// Calculate window size

if windowSize == k {

    // Process the complete window

    // Remove element leaving the window

    // Move left forward
}


}

Time Complexity: O(n)

Space Complexity: O(1)
*/

package main

import (
	"fmt"
)

func findMaxAverage(nums []int, k int) float64 {

	left := 0
	sum := 0
	var maxAvg float64

	for right := 0; right < len(nums); right++ {

		windowSize := right - left + 1
		sum += nums[right]

		if windowSize == k {

			avg := float64(sum) / float64(k)

			if left == 0 {
				maxAvg = avg
			}
			if avg > maxAvg {
				maxAvg = avg
			}
			sum -= nums[left]
			left++
		}
	}
	return maxAvg
}

func main() {

	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4
	result := findMaxAverage(nums, k)
	fmt.Println(result)

}
