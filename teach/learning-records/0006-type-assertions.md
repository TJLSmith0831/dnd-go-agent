---
name: type-assertions
description: Lesson 06 — type assertions, two-value safe form, type switches, named types as lightweight enums. Score 4/4.
metadata:
  type: feedback
---

## What was covered
- Single-value assertion `v := x.(T)` — panics if wrong type
- Two-value assertion `v, ok := x.(T)` — safe, ok is false on mismatch
- Type switch `switch v := x.(type)` — v is typed as the matched concrete type in each case
- `default` case in type switch holds the original interface type
- Named types: `type ZeroHPOutcome string` creates a distinct type the compiler enforces
- Design smell: asserting to call a method that all interface implementors have → add it to the interface instead
- Legitimate assertion: behavior genuinely specific to one concrete type (death saves = PC only)

## Score
4/4 — sixth consecutive perfect score

## Zone of proximal development
Ready for: D&D spellcasting (overdue for a D&D lesson after four straight Go lessons), then goroutines/channels.
