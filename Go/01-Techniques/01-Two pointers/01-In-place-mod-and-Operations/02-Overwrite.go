/*
Technique: In-place Modification

Pattern: None (Technique Practice)

Operation: Overwrite

Objective:
Learn how to compact valid elements toward the beginning of the same array
without using an extra array.

Key Learning:
Use a write index to track the next position where a valid element should be
placed. Traverse the array, and whenever a valid element is found, overwrite
the element at the write position and increment the write index. Invalid
elements are ignored, allowing future valid elements to naturally overwrite
them.

Important Concept:
This operation performs a copy, NOT a swap. The destination value is not
preserved because it is considered unnecessary.

Time Complexity: O(n)
Space Complexity: O(1)
*/

package main

func main() {

	arr := [5]int{5, 0, 8, 0, 3}

	write := 0

	for _, value := range arr {
		if value != 0 {
			arr[write] = value
			write++
		}
	}

}
