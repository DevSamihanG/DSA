/*
Technique: In-place Modification

Pattern: None (Technique Practice)

Operation: Swap

Objective:
Learn how to exchange two elements in an array without losing either value.

Key Learning:
Before overwriting a value, preserve it in a temporary variable.
Once the original value is safe, replace the first element and then restore the saved value.

Time Complexity: O(1)
Space Complexity: O(1)
*/

package main

import (
	"fmt"
)

func main() {

	arr := [4]int{10, 20, 30, 40}

	temp := arr[1]

	arr[1] = arr[3]

	arr[3] = temp
	fmt.Println(arr)

}
