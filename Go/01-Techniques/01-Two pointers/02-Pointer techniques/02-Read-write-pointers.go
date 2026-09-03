/*
Technique: Pointer Technique

Pattern: None (Technique Practice)

Operation: Read / Write Pointers

Objective:
Learn how two pointers moving in the same direction can be used to
read elements from an array and write selected elements toward the
beginning of the same array.

Key Learning:
The read pointer scans every element of the array.
The write pointer tracks the next position where a valid element
should be placed.

When a valid element is found:
1. Copy the element from the read position to the write position.
2. Move the write pointer forward.

The read pointer and write pointer do NOT necessarily move together.
The read pointer always scans the array, while the write pointer
moves only when an element needs to be written.

Important Concept:
Read = "What element am I looking at?"
Write = "Where should the next valid element go?"

Time Complexity: O(n)
Space Complexity: O(1)
*/

package main

import "fmt"

func main() {

	arr := [5]int{5, 0, 8, 0, 3}

	write := 0

	for read := 0; read < len(arr); read++ {

		if arr[read] != 0 {
			arr[write] = arr[read]
			write++
		}
	}

	fmt.Println(arr)
}
