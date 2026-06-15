---
name: conditions
description: Lesson 12 — D&D 5e conditions (8 key conditions, cascade rules, advantage/disadvantage effects) + Go iota bit flags and ConditionSet.
metadata:
  type: feedback
---

## What was covered
- 8 conditions: Blinded, Frightened, Incapacitated, Paralyzed, Poisoned, Prone, Restrained, Stunned
- Cascade: Paralyzed and Stunned include all effects of Incapacitated
- Prone: melee attackers get advantage, ranged attackers get disadvantage
- Resolution path with conditions: check attacker disadvantage, check defender grants advantage, cancel if both
- `iota` auto-increments from 0 in each const block
- `1 << iota` gives each constant a unique bit (powers of two: 1, 2, 4, 8…)
- `|=` to apply, `&^=` to remove (AND-NOT / bit-clear), `&` to check
- `ConditionSet` as a `uint16` wrapping struct — zero value is correct (no conditions)
- Go's `&^` operator = Python's `&= ~x`

## TDD exercise
Add `pkg/combat/conditions.go`: `Condition uint16` named type, 8 constants with `1 << iota`, `ConditionSet` with `Apply`/`Remove`/`Has`, `AttackDisadvantage()`, `GrantsAdvantageToAttackers()`. Five tests in `conditions_test.go`.

## Next wiring step
Add `Conditions ConditionSet` to `Monster` and `PlayerCharacter`, update `ResolveAttack` to query conditions before rolling, implement rolling with advantage/disadvantage (roll twice, take higher/lower).

## Quiz fix applied
All four answer options written to equal length — no formatting clue about correct answer.

## Zone of proximal development
Ready for: Lesson 13 — Postgres & pgx. Connect to a real database, write campaign state, query it. The Narrator agent needs persistent world memory.
