Pattern 1: Two Pointers

1. Theory & Recognition  ← Next
2. Opposite Direction
3. Same Direction
4. Classic Problems
5. Mixed Problems

# Two Pointer Pattern

## Definition

The Two Pointer Pattern is an algorithmic pattern where two indices (pointers)
are used to traverse or process a data structure together instead of using a
single index.

The pointers move according to specific rules depending on the problem, allowing
us to solve many problems in O(n) time instead of using nested loops.

---

## Why do we use Two Pointers?

A single pointer can only observe one position at a time.

Two pointers allow us to:

- Compare two elements.
- Search from both ends.
- Process two positions simultaneously.
- Avoid unnecessary repeated traversals.
- Reduce O(n²) solutions to O(n) in many cases.

---

## Core Idea

The algorithm is NOT about having two variables.

The algorithm is about defining movement rules for two pointers.

For every problem, ask:

1. Where does the left pointer start?
2. Where does the right pointer start?
3. When does the left pointer move?
4. When does the right pointer move?
5. When should the algorithm stop?

If you can answer these five questions, you've almost solved the problem.

---

## Common Pointer Movements

### 1. Opposite Direction

L →           ← R

Used when processing both ends of a structure.

Examples:
- Reverse Array
- Valid Palindrome
- Two Sum II (Sorted Array)
- Container With Most Water

---

### 2. Same Direction

L →
R →

Both pointers move forward.

One pointer usually explores while the other marks or maintains a position.

Examples:
- Remove Duplicates
- Move Zeroes
- Merge Sorted Arrays

---

### 3. Sliding Window

L → →
R → →

Both pointers move forward independently.

The distance between them changes depending on the problem.

Examples:
- Longest Substring Without Repeating Characters
- Maximum Sum Subarray of Size K
- Minimum Window Substring

---

## Recognition Checklist

Think about Two Pointers if you notice:

✓ The input is an array or string.
✓ You need to compare two positions.
✓ The problem mentions sorted data.
✓ You need to process elements from both ends.
✓ You want to avoid nested loops.
✓ You need an O(n) solution.

---

## Time Complexity

Most Two Pointer algorithms run in O(n).

Reason:

Even though there are two pointers, each pointer usually moves only in one
direction and never revisits previous positions.

---

## Key Learning

Two Pointers is a Pattern.

Swap, Shift, Overwrite and Reverse are Techniques.

The pattern decides HOW to explore the data.

The techniques decide WHAT operation to perform once the correct positions are
found.


## Problem Solving Framework

Whenever you encounter a Two Pointer problem, answer these questions before
writing any code.

1. Which Two Pointer variant is this?
   - Opposite Direction
   - Same Direction
   - Sliding Window

2. Where do the pointers start?

3. What operation happens every iteration?
   - Compare
   - Swap
   - Overwrite
   - Add
   - Skip
   - Something else

4. When do the pointers move?

5. When does the algorithm stop?


## Problems Covered

### Opposite Direction

1. Reverse Array
   Technique Used: Swap

2. Palindrome Check
   Technique Used: Compare


## Pointer Movement Rulebook

Every Two Pointer problem can be understood by answering one question:

**How are the pointers allowed to move?**

There are three fundamental movement types.

---

### Movement Type 1: Fixed Movement

Both pointers move every iteration.

Pattern:

L →       ← R

Rules:

- Left always moves right.
- Right always moves left.
- No decisions are made.
- Movement is predetermined.

Examples:

✓ Reverse Array
✓ Palindrome Check

---

### Movement Type 2: Decision-Based Movement

Only one pointer moves depending on the current comparison.

Pattern:

L →       ← R

Rules:

- Compare the current values.
- Decide which pointer should move.
- The other pointer remains in place.
- Pointer movement depends on the problem.

Examples:

✓ Two Sum II
✓ Container With Most Water

---

### Movement Type 3: Independent Movement

Both pointers move in the same direction.

Pattern:

L →
R →

Rules:

- One pointer explores.
- One pointer maintains information.
- The pointers move independently.
- They do not necessarily move together.

Examples:

✓ Remove Duplicates
✓ Move Zeroes
✓ Sliding Window (advanced)



## Decision-Based Movement Rule

Current Sum < Target
→ Need a bigger sum.
→ Move the left pointer.

Current Sum > Target
→ Need a smaller sum.
→ Move the right pointer.

Reason:
In a sorted array:
- Moving left to the right increases the chosen left value.
- Moving right to the left decreases the chosen right value.