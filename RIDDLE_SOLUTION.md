# Logic Riddle Solution

## The Riddle

Mister X knows the sum of numbers A and B. Mister Y knows the product of numbers A and B.

**Mister Y says to Mister X:** "I don't know what A and B are."

**Mister X then says:** "I know you don't know."

**Mister Y then says:** "In that case, I do know what A and B are."

**Question:** What are A and B?

**Constraints:** Both numbers are integers between 2 and 99 (inclusive).

## The Answer

**A = 4, B = 13**
- Sum = 17
- Product = 52

## Logical Deduction

### Step 1: Y says "I don't know"
Product = 52 can be factored as:
- 2 × 26
- 4 × 13

Since there are multiple factorizations, Y cannot determine A and B uniquely.

### Step 2: X says "I know you don't know"
Sum = 17. X knows that for ANY pair of numbers that sum to 17, their product must have multiple factorizations.

All pairs that sum to 17:
- (2,15) → 30 = 2×15 = 3×10 = 5×6 ✓
- (3,14) → 42 = 2×21 = 3×14 = 6×7 ✓
- (4,13) → 52 = 2×26 = 4×13 ✓
- (5,12) → 60 = 2×30 = 3×20 = 4×15 = 5×12 = 6×10 ✓
- (6,11) → 66 = 2×33 = 3×22 = 6×11 ✓
- (7,10) → 70 = 2×35 = 5×14 = 7×10 ✓
- (8,9) → 72 = 2×36 = 3×24 = 4×18 = 6×12 = 8×9 ✓

All products are ambiguous! Therefore, X can confidently say "I know you don't know."

The sums with this special property are: [11, 17, 23, 27, 29, 35, 37, 41, 47, 53]

### Step 3: Y says "Now I know"
After hearing X's statement, Y knows the sum must be one of these special sums.

For product 52, the factorizations are:
- (2, 26) → sum = 28 (NOT in the special list)
- (4, 13) → sum = 17 (YES, in the special list!)

Only one factorization has a sum in the special list, so Y can now uniquely determine: **A = 4, B = 13**

### Step 4: Verification that X also knows
After Y says "Now I know," X can determine which pair from sum 17 is the answer, because only the pair (4, 13) leads to a product where Y could deduce the answer.

## Running the Solution

```bash
go run riddle_solver.go
```

This will output the logical deduction process and verify the solution.
