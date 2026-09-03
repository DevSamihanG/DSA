# 04 — Add / Remove Window Elements

This is essentially the **general rule behind Window State Management**.

So far, we have seen three specific examples:

```text
Running Sum:
Entering element → add its value
Leaving element  → subtract its value

Running Count:
Entering element → check condition → count++
Leaving element  → check condition → count--

Frequency Map:
Entering element → frequency[element]++
Leaving element  → frequency[element]--
```

Now let's step back and identify the **common pattern**.

## The general idea

Whenever the window changes, an element either:

```text
ENTERS the window
```

or:

```text
LEAVES the window
```

Your job is to update whatever **state** you are maintaining.

```text
                 WINDOW STATE
                      │
        ┌─────────────┴─────────────┐
        │                           │
Element enters                 Element leaves
        │                           │
   Add its effect              Remove its effect
```

For example, suppose:

```text
Window: [10, 20, 30]
```

and you're maintaining a sum:

```text
State = sum = 60
```

When `40` enters:

```text
[10, 20, 30, 40]
```

Update the state by including `40`:

```text
sum += 40
```

When `10` leaves:

```text
[20, 30, 40]
```

Update the state by removing `10`:

```text
sum -= 10
```

The same principle works for every type of state:

| State being maintained | Element enters               | Element leaves               |
| ---------------------- | ---------------------------- | ---------------------------- |
| Sum                    | Add value                    | Subtract value               |
| Count of even numbers  | `count++` if even            | `count--` if even            |
| Frequency map          | Increase frequency           | Decrease frequency           |
| Maximum/minimum        | May require special handling | May require special handling |

### The core mental model

> **A sliding window does not usually rebuild its state from scratch. When the window changes, update the existing state according to the element entering or leaving.**

This is the concept that connects everything we just practiced:

```text
Window changes
     ↓
Identify what changed
     ↓
Which element entered?
Which element left?
     ↓
Update only the affected state
```

For this particular subtopic, we **don't need another complicated new program**. You have already implemented the add/remove behavior using:

* Running Sum
* Running Count
* Frequency Map

So I would treat this as a **conceptual supporting technique** rather than creating another redundant code file.

At this point, your structure becomes:

```text
02-Window-State-Management
│
├── 01-Running-Sum          ✅
├── 02-Running-Count        ✅
├── 03-Frequency-Map        ✅
└── 04-Add-Remove-Elements  ✅
```

The next major thing we should learn is **how these fundamentals combine into an actual Sliding Window pattern**—starting with the distinction between **fixed-size and variable-size sliding windows**.
