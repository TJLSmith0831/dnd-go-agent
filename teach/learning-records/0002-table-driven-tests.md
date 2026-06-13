---
name: table-driven-tests
description: Lesson 02 — Go testing, table-driven pattern, t.Run subtests, blank identifier. Score 4/4.
metadata:
  type: feedback
---

## What was covered
- `*_test.go` naming convention; compiled only during `go test`
- `testing.T` — `t.Error`/`t.Errorf` (fail and continue) vs `t.Fatal`/`t.Fatalf` (fail and stop)
- Table-driven tests: inline anonymous struct slice, range over cases
- `t.Run(name, func)` for named subtests; runnable individually with `-run TestFoo/subtest_name`
- Blank identifier `_` — explicit discard of a return value; requires conscious choice
- Testing random output by contract (range check × 1000) rather than fixed seed
- `for range N` syntax (Go 1.22+)

## Score
4/4 — second consecutive perfect score

## Zone of proximal development
Moving fast. Ready for: D&D combat mechanics (action economy, attack rolls, AC) — due for a D&D lesson after two Go lessons. After that: Go methods on structs + interfaces.

## Notes
User is writing code alongside lessons without being prompted. Good sign — apply knowledge immediately while lesson is fresh.
