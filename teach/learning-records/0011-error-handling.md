---
name: error-handling
description: Lesson 11 — Go error handling: sentinel errors, fmt.Errorf %w wrapping, errors.Is, errors.As, custom error types with Unwrap.
metadata:
  type: feedback
---

## What was covered
- `error` is an interface: any type with `Error() string` satisfies it
- Sentinel errors: `var ErrX = errors.New("...")` — package-level, express a kind of failure
- `fmt.Errorf("context: %w", err)` — wraps an error, preserves the chain
- `fmt.Errorf("context: %s", err)` — embeds the string only, severs the chain
- `errors.Is(err, target)` — checks identity through the entire unwrap chain
- `errors.As(err, &target)` — extracts a typed error from the chain, populates target's fields
- Custom error types: struct implementing `Error() string` + `Unwrap() error` for chain traversal
- `Unwrap()` is what lets `errors.Is` and `errors.As` see through a custom type to the sentinel it wraps

## TDD exercise
Build `pkg/referee`: `Skill` named type, `ErrUnknownSkill` sentinel, `RefereeError` struct with `Input`/`Reason` fields and `Unwrap()`, `ResolveSkill(input string) (Skill, error)` with keyword matching. Three tests: recognises stealth, returns `ErrUnknownSkill`, `errors.As` extracts input field.

## Key design note
Keyword matching in `ResolveSkill` is a placeholder. In the real bot, the LLM resolves intent — the error types and function signature stay the same, only the matching logic changes.

## Zone of proximal development
Ready for: Lesson 12 — D&D conditions (Poisoned, Frightened, Paralyzed…). Last rules layer before the full combat loop is complete.
