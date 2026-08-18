/*
Container With Most Water

Core Idea:
Two walls and the x-axis form a container. The goal is to find the pair
of walls that can hold the maximum amount of water.

Parameters:
- Width = right - left
- Limiting Height = min(height[left], height[right])
- Current Area = Width × Limiting Height

Algorithm:
1. Calculate the current width.
2. Find the limiting (shorter) wall.
3. Compute the current area.
4. Update the maximum area.
5. Move the pointer at the shorter wall.
6. Repeat until the pointers meet.

Movement Rule:
Always move the pointer at the shorter wall.

Reason:
The shorter wall limits the water level.
Moving the taller wall only reduces the width and cannot increase the
current limiting height. Only moving the shorter wall gives a chance of
finding a taller wall and producing a larger area.

Pattern:
Two Pointers (Opposite Direction)

Time: O(n)
Space: O(1)
*/

package main

import (
	"fmt"
)

func main() {

	//0  1  2  3  4  5  6  7  8
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	left := 0
	right := len(height) - 1
	maxArea := 0

	//width = 7
	//lim height = min(1, 7) = 1
	//currentArea = width * height
	currArea := 0

	for left < right {
		width := right - left
		if height[left] < height[right] {
			currArea = width * height[left]
		} else {
			currArea = width * height[right]
		}
		if currArea > maxArea {
			maxArea = currArea
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	fmt.Println(maxArea)
}
