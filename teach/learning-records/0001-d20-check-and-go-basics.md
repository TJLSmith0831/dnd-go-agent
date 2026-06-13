---
name: d20-check-and-go-basics
description: Lesson 01 — d20 ability check mechanics + core Go syntax. Score 4/4.
metadata:
  type: feedback
---

## What was covered
- D&D 5e ability check: d20 + modifier ≥ DC → success
- Modifier formula: ⌊(score − 10) / 2⌋
- Standard DC table (5, 10, 15, 20, 25, 30)
- Go: package declarations, imports (unused = compile error)
- Go: capitalized name = exported; lowercase = package-private
- Go: floor division vs truncation (`math.Floor` needed for correct modifiers on odd scores below 10)
- Go: errors as return values (no exceptions), `(result, error)` tuple pattern
- Go: structs as the alternative to classes; zero values
- Go: `:=` for declare-and-assign

## Score
4/4 — all questions correct

## Zone of proximal development
Ready for: Go module setup, `go test`, table-driven tests. Can skip "what is a package" explanations. Introduce Go idioms at pace.

## Key insight to revisit
Floor vs truncation division is the first concrete place where Go behavior diverges from Python in a semantically meaningful way. Worth confirming it sticks once they write the actual test.
