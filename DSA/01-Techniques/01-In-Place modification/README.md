# Technique 01 - In-place Modification

## Definition

**In-place Modification** is a technique where we solve a problem by modifying the original data structure instead of creating a new one.

The primary goal is to reduce extra space usage while efficiently updating the existing data.

---

# Why does this technique exist?

Suppose we have the following array:

```text
[4, 2, 9, 1, 7]
```

Now we are asked to remove every even number.

There are two approaches.

### Approach 1 - Create a New Array

```text
Original Array
      ↓
Create New Array
      ↓
Copy only the required elements
```

This requires additional memory.

---

### Approach 2 - Modify the Existing Array

```text
Original Array
      ↓
Shift / Overwrite Elements
      ↓
Result
```

Instead of allocating another array, we modify the existing one.

This approach is called **In-place Modification**.

---

# Purpose

The purpose of this technique is to modify the existing data structure while using little or no extra memory.

In interviews, this usually means reducing the space complexity from:

```text
O(n)
      ↓
O(1)
```

---

# Recognition Clues

Whenever a problem contains phrases such as:

- Modify the array
- Rearrange the elements
- Remove elements
- Move elements
- Rotate the array
- Reverse the array
- Partition the array
- Solve using constant extra space

It is a strong indication that **In-place Modification** might be required.

---

# Data Structures

This technique is commonly used with:

- Arrays
- Strings (in some languages)
- Linked Lists

---

# Core Operations

Almost every in-place algorithm is built using one or more of the following operations.

## 1. Swap

Exchange two elements.

Example:

```text
Before

1 2 3

After

2 1 3
```

---

## 2. Overwrite

Replace one element with another.

Example:

```text
Before

1 0 3

After

1 3 3
```

Notice that we copied a value over another value instead of swapping.

---

## 3. Shift

Move a sequence of elements to create or remove space.

Example:

```text
Before

1 2 3 4

After

2 3 4 _
```

---

## 4. Reverse

Rearrange elements in the opposite order.

Example:

```text
Before

1 2 3 4

After

4 3 2 1
```

---

# Mental Model

Imagine arranging books on a bookshelf.

Current shelf:

```text
A B C D E
```

Suppose book **B** is removed.

You do not buy another shelf.

Instead, you slide every remaining book to the left.

```text
A C D E
```

The same shelf is being modified.

This is exactly what **In-place Modification** means.

---

# Why are we learning this before Two Pointers?

Two Pointers answers the question:

> Which indices should move?

In-place Modification answers the question:

> What should happen to the data once those indices move?

Example:

Move Zeroes

Two Pointers decides:

```text
Left Pointer

Right Pointer
```

In-place Modification decides:

```text
Swap

Overwrite

Update
```

Knowing the pattern alone is not enough.

The technique tells us **how** to modify the array.

---

# Summary

- In-place Modification is a **technique**, not a pattern.
- It focuses on modifying the existing data structure instead of creating a new one.
- The main objective is to reduce extra space complexity.
- The four fundamental operations are:
  - Swap
  - Overwrite
  - Shift
  - Reverse
- Many patterns, especially **Two Pointers**, heavily rely on this technique.