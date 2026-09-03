package main
/*
Technique: In-Place Modification

Pattern: None (Technique Practice)

Operation: Filling Remaining Positions

Objective:
Learn how to fill the unused portion of an array after valid elements
have been compacted toward the beginning.

Key Learning:
After an overwrite operation, the write index tells us where the
valid portion of the array ends.

The positions from write to the end of the array are the remaining
positions.

These positions can then be filled with a required value.

Important Concept:
This is commonly a second phase after an overwrite operation:

Phase 1:
Compact valid elements toward the beginning.

Phase 2:
Fill the remaining positions with the required value.

The write index acts as the boundary between the valid portion and
the remaining portion.

Example:

[0, 1, 0, 3, 12]

After compaction:

[1, 3, 12, _, _]
             ↑
           write

Fill the remaining positions:

[1, 3, 12, 0, 0]

Time Complexity: O(n)
Space Complexity: O(1)
*/

package main

import "fmt"

func main() {

	arr := [5]int{0, 1, 0, 3, 12}

	write := 0

	// Phase 1: Compact valid elements.
	for read := 0; read < len(arr); read++ {

		if arr[read] != 0 {
			arr[write] = arr[read]
			write++
		}
	}

	// Phase 2: Fill remaining positions.
	for i := write; i < len(arr); i++ {
		arr[i] = 0
	}

	fmt.Println(arr)
}