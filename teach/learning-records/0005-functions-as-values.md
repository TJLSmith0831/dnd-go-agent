---
name: functions-as-values
description: Lesson 05 — functions as first-class values, closures, sort.Slice, make(). Score 4/4.
metadata:
  type: feedback
---

## What was covered
- Function types: written as `func(paramTypes) returnType`
- Assigning, passing, and returning functions as values
- Closures: capture variables from surrounding scope by reference, not by value
- `sort.Slice(slice, func(i, j int) bool)` — less function as a closure over the slice
- Descending sort: flip `<` to `>`
- `make([]T, n)` — pre-allocate a slice of length n vs `[]T{}` (empty, must append)
- Closure gotcha: captures variable reference not value — Go 1.22+ fixes it for loop vars specifically

## Score
4/4 — fifth consecutive perfect score

## Applied immediately
Wrote `TestRollInitiative` test first (red), then updated `RunInitiative` to return `([]InitiativeEntry, error)` sorted descending. Followed TDD pattern from the lesson correctly.

## Zone of proximal development
Ready for: type assertions and type switches — needed immediately for combat loop (PC vs monster at 0 HP). Then goroutines/channels.
