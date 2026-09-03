/*
Maximum Sum Subarray of Size K

Pattern:
Fixed-Size Sliding Window

Idea:
Use a sliding window of size k.

* right expands the window.
* sum maintains the sum of the current window.
* maxSum stores the largest window sum found so far.
* left moves forward after processing a complete window.

Algorithm:

1. Start left at index 0.
2. Move right through the array.
3. Add arr[right] to sum as each element enters the window.
4. Calculate the current window size.
5. When windowSize == k:

   * Compare sum with maxSum.
   * Update maxSum if the current sum is larger.
   * Remove arr[left] from sum because it is leaving the window.
   * Move left forward to slide the window.
6. Return maxSum.

Reason:
The window is maintained incrementally instead of calculating
the sum from scratch for every window.

When a new element enters:

```
sum += arr[right]
```

When an old element leaves:

```
sum -= arr[left]
```

This allows us to move the window forward efficiently while
maintaining its sum.

Time Complexity: O(n)

Space Complexity: O(1)
*/

package main

import (
	"fmt"
)

func maxSum(arr [6]int, k int) int {

	left := 0
	sum := 0
	maxSum := 0

	for right := 0; right < len(arr); right++ {

		windowSize := right - left + 1
		sum += arr[right]

		if windowSize == k {
			if sum > maxSum {
				maxSum = sum
			}
			sum -= arr[left]
			left++
		}

	}
	return maxSum
}

func main() {

	arr := [6]int{2, 1, 5, 1, 3, 2}
	k := 3
	result := maxSum(arr, k)
	fmt.Println(result)

}
