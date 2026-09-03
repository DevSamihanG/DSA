Absolutely. Since we have now completed the **Sliding Window fundamentals and state-management techniques**, this is the right time to create the main **Sliding Window Pattern** write-up.

I would keep it parallel to your Two Pointer documentation, but make one important distinction clear:

> **Sliding Window uses two boundaries, but its main purpose is to maintain and process a contiguous range efficiently.**

Here is the write-up I recommend:

# Pattern 2: Sliding Window

## 1. Theory & Recognition

**# Sliding Window Pattern**

## Definition

The Sliding Window Pattern is an algorithmic pattern used to process a **contiguous range of elements** in an array or string.

Two boundaries, usually called `left` and `right`, define the current window.

```text
left                     right
  ↓                        ↓

[10, 20, 30, 40, 50, 60]
     [20, 30, 40]
```

The window represents all elements from:

```text
arr[left] → arr[right]
```

including every element between them.

The boundaries move through the data according to rules defined by the problem.

As the window changes:

* New elements may enter the window.
* Existing elements may leave the window.
* The state of the window is updated efficiently.

The goal is to avoid repeatedly processing the same range from scratch.

---

## Why do we use Sliding Window?

Consider a problem asking us to process every contiguous subarray.

A brute-force approach may repeatedly calculate information for overlapping ranges:

```text
[10, 20, 30]
    [20, 30, 40]
        [30, 40, 50]
```

Notice that many elements are processed repeatedly.

For example:

```text
10 + 20 + 30
20 + 30 + 40
30 + 40 + 50
```

The Sliding Window approach reuses information from the previous window.

Instead of recalculating:

```text
20 + 30 + 40
```

from scratch, we can:

```text
Previous sum = 60

Remove 10
Add 40

New sum = 90
```

Sliding Window helps us:

* Process contiguous ranges efficiently.
* Reuse information from the previous window.
* Avoid unnecessary repeated calculations.
* Maintain window state dynamically.
* Often reduce O(n²) approaches to O(n).

---

## Core Idea

The Sliding Window Pattern is **not simply about using `left` and `right` pointers**.

The core idea is:

> **Maintain a contiguous range and update its state as elements enter and leave the range.**

The general process is:

```text
Define the window
        ↓
Process the current window
        ↓
Move a boundary
        ↓
An element enters or leaves
        ↓
Update the window state
        ↓
Process the new window
```

For every Sliding Window problem, ask:

1. What does the current window represent?
2. Where does the window start?
3. How does the window expand?
4. When should the window shrink?
5. What information must be maintained about the window?
6. When should the answer be updated?
7. When does the algorithm stop?

If you can answer these questions, the Sliding Window solution becomes much easier to construct.

---

# Window Fundamentals

A Sliding Window always represents a **contiguous range**.

For:

```text
arr = [10, 20, 30, 40, 50, 60]

left = 1
right = 3
```

The current window is:

```text
[20, 30, 40]
```

because it contains:

```text
arr[1], arr[2], arr[3]
```

The window cannot skip elements.

This:

```text
[20, 30, 40]
```

is a valid window.

This:

```text
[20, 40, 50]
```

is not a contiguous window because `30` is skipped.

---

## Window Boundaries

The two boundaries have specific responsibilities:

```text
left  → starting position of the window
right → ending position of the window
```

Example:

```text
Index:   0    1    2    3    4    5
Value:  10   20   30   40   50   60
              ↑         ↑
            left       right
```

The window includes everything between them.

```text
[20, 30, 40]
```

The boundaries do **not** represent only two selected elements.

They define the entire range.

---

# Window Movement

The window changes by moving its boundaries.

## 1. Expand Window

To expand the window:

```text
right++
```

Example:

```text
Before:

[20, 30, 40]

After right++:

[20, 30, 40, 50]
```

A new element enters the window:

```text
arr[right]
```

General idea:

```text
Move right
    ↓
New element enters
    ↓
Update window state
```

---

## 2. Shrink Window

To shrink the window:

```text
left++
```

Example:

```text
Before:

[20, 30, 40, 50]

After left++:

[30, 40, 50]
```

The element at the **old `left` position** leaves the window.

General idea:

```text
Remove old left element
        ↓
Move left
```

This ordering is important when maintaining window state.

For example:

```text
frequency[arr[left]]--
left++
```

The old element must be processed before `left` moves.

---

# Window State Management

A window often needs to maintain information about its current elements.

This information is called the **window state**.

Examples include:

```text
Sum
Count
Frequency Map
Distinct Element Count
Maximum
Minimum
Other problem-specific information
```

The important principle is:

> **Do not rebuild the state from scratch every time the window changes. Update it based on what entered or left.**

---

## Running Sum

Suppose:

```text
Window: [10, 20, 30]

sum = 60
```

If `40` enters:

```text
right++
sum += arr[right]
```

The new state becomes:

```text
Window: [10, 20, 30, 40]

sum = 100
```

If `10` leaves:

```text
sum -= arr[left]
left++
```

The new state becomes:

```text
Window: [20, 30, 40]

sum = 90
```

---

## Running Count

Suppose we count even numbers.

```text
Window: [10, 3, 8]

count = 2
```

When an element enters:

```text
If it satisfies the condition:
    count++
```

When an element leaves:

```text
If it satisfies the condition:
    count--
```

Only elements satisfying the required condition affect the state.

---

## Frequency Map

A frequency map stores:

```text
element → number of occurrences
```

Example:

```text
Window:

[1, 2, 1, 3]

Frequency:

1 → 2
2 → 1
3 → 1
```

When an element enters:

```text
frequency[element]++
```

When an element leaves:

```text
frequency[element]--
```

If its frequency becomes zero:

```text
delete(frequency, element)
```

This keeps the map synchronized with the elements currently inside the window.

---

# Add / Remove Window Elements

This is the general rule behind Sliding Window state management.

Whenever the window changes, identify:

```text
What entered?
What left?
```

Then update the state accordingly.

```text
ELEMENT ENTERS
      ↓
Add its contribution to the state


ELEMENT LEAVES
      ↓
Remove its contribution from the state
```

Examples:

| Window State      | Element Enters    | Element Leaves               |
| ----------------- | ----------------- | ---------------------------- |
| Sum               | Add value         | Subtract value               |
| Conditional count | Check → `count++` | Check → `count--`            |
| Frequency map     | `frequency[x]++`  | `frequency[x]--`             |
| Distinct elements | May increase      | May decrease                 |
| Maximum/Minimum   | Problem-dependent | May require special handling |

---

# Fixed-Size vs Variable-Size Window

There are two major types of Sliding Window problems.

## 1. Fixed-Size Window

The window size is predetermined.

Example:

> Find the maximum sum of every subarray of size `k`.

If:

```text
k = 3
```

then every window must contain exactly three elements.

```text
[10, 20, 30]
    [20, 30, 40]
        [30, 40, 50]
```

The process is generally:

```text
Expand
    ↓
Window reaches size k
    ↓
Process window
    ↓
Shrink from left
    ↓
Continue
```

The window repeatedly:

```text
expands → processes → shrinks
```

while maintaining the same size.

---

## 2. Variable-Size Window

The window size is determined by a condition.

Example:

> Find the longest subarray containing at most `k` distinct elements.

The window may grow:

```text
[1]
[1, 2]
[1, 2, 1]
[1, 2, 1, 3]
```

If the condition becomes invalid, shrink the window until it becomes valid again.

General process:

```text
Expand window
      ↓
Check condition
      ↓
Valid?
 ┌────┴────┐
Yes        No
↓          ↓
Continue   Shrink window
           until valid
```

The size is therefore not fixed.

---

# Fixed-Size vs Variable-Size

| Fixed-Size Window                                     | Variable-Size Window                     |
| ----------------------------------------------------- | ---------------------------------------- |
| Window size is predetermined                          | Window size changes based on a condition |
| Example: size `k`                                     | Example: at most `k` distinct elements   |
| Usually expand and then shrink at a predictable point | Shrink only when required                |
| Every valid window has the same size                  | Valid windows can have different sizes   |

---

# Recognition Checklist

Think about Sliding Window when you notice:

✓ The input is an array or string.

✓ The problem deals with a **contiguous subarray or substring**.

✓ The problem asks for information about a range.

✓ The range can move through the input.

✓ Adjacent ranges overlap significantly.

✓ You are repeatedly calculating similar information.

✓ The problem asks for longest, shortest, maximum, minimum, count, or sum of a contiguous range.

✓ A brute-force solution would examine many overlapping subarrays or substrings.

A particularly strong signal is language such as:

```text
Longest substring
Shortest subarray
Maximum sum subarray
Minimum length subarray
Subarray of size K
Substring containing...
At most K distinct...
Without repeating...
```

---

# When NOT to Use Sliding Window

Sliding Window is specifically designed for **contiguous ranges**.

If the problem allows arbitrary elements to be selected:

```text
[10, 30, 60]
```

while skipping elements between them, that is generally **not a normal Sliding Window problem**.

Remember:

> **Sliding Window = contiguous range.**

Also, not every problem with `left` and `right` pointers is a Sliding Window problem.

For example:

```text
Two Sum II
```

uses two pointers, but they represent two selected positions:

```text
left         right
  ↓             ↓

[2, 7, 11, 15]
```

They do not define a range whose internal state is being maintained.

That is **Opposite Direction Two Pointers**, not Sliding Window.

---

# Time Complexity

Most standard Sliding Window algorithms run in:

```text
O(n)
```

The important reason is that the boundaries generally move only forward.

```text
left  → → → →
right → → → →
```

Even though there may be nested loops, such as:

```text
for right moves forward {
    for condition is invalid {
        left moves forward
    }
}
```

this does not automatically mean O(n²).

Each element is generally:

* Added to the window once.
* Removed from the window once.

Therefore, the total amount of boundary movement is often proportional to `n`.

However, complexity depends on the state operation. If updating or calculating window state itself takes significant time, the total complexity can be higher.

---

# Key Learning

Sliding Window is a **Pattern**.

The techniques we practiced earlier support that pattern:

```text
Contiguous Range
Left / Right Boundaries
Expand Window
Shrink Window
Running Sum
Running Count
Frequency Map
Add / Remove Elements
```

The Sliding Window pattern decides:

> **How to maintain and move through a contiguous range.**

The supporting techniques decide:

> **How to maintain information about that range.**

For example:

```text
Pattern:
Sliding Window

Window movement:
Expand / Shrink

State:
Frequency Map

Operation:
Add entering character
Remove leaving character
```

These pieces work together.

---

# Problem Solving Framework

Whenever you encounter a Sliding Window problem, answer these questions before writing code:

### 1. Is the problem asking about a contiguous range?

```text
Subarray?
Substring?
Consecutive elements?
```

If yes, Sliding Window may apply.

### 2. What type of window is it?

```text
Fixed-size?
Variable-size?
```

### 3. What does `left` represent?

Usually:

```text
Start of current window
```

### 4. What does `right` represent?

Usually:

```text
End of current window / exploring boundary
```

### 5. What happens when `right` moves?

```text
Which element enters the window?
```

Update the state.

### 6. What condition determines shrinking?

For fixed-size:

```text
Window became too large.
```

For variable-size:

```text
The required condition became invalid.
```

### 7. What happens when `left` moves?

```text
Which element leaves the window?
```

Remove its contribution from the state.

### 8. When should the answer be updated?

This is problem-dependent.

For example:

```text
When window size = k
```

or:

```text
When window is valid
```

### 9. When does the algorithm stop?

Usually:

```text
right reaches the end of the input.
```

---

# The Sliding Window Rulebook

Every Sliding Window problem can be reduced to this flow:

```text
1. Expand the window
        ↓
2. Add the entering element to the window state
        ↓
3. Check the window condition
        ↓
4. If necessary, shrink the window
        ↓
5. Remove the leaving element from the window state
        ↓
6. Update the answer when appropriate
        ↓
7. Continue
```

The most important question throughout the problem is:

> **What entered the window, what left the window, and what state must I update?**

That is the central mental model of the **Sliding Window Pattern**.

