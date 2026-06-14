---
name: methods-and-interfaces
description: Lesson 04 — Go methods on structs, value vs pointer receivers, implicit interface satisfaction, fmt.Stringer. Score 4/4.
metadata:
  type: feedback
---

## What was covered
- Method syntax: receiver comes between `func` and the method name
- Value receiver: gets a copy, reads only
- Pointer receiver: gets the original, can mutate state
- Consistency rule: if any method uses a pointer receiver, all should
- Interface satisfaction is implicit — no `implements` keyword, no declaration
- Interface is defined by the consumer, not the types that satisfy it
- `fmt.Stringer`: `String() string` makes a type auto-format with `fmt` functions
- Pointer receiver gotcha: `*Monster` satisfies the interface, `Monster` does not
- `switch` without condition (used in DeathSavingThrow)

## Score
4/4 — fourth consecutive perfect score

## Applied immediately
User implemented `DeathSavingThrow`, `RunInitiative`, and is building `ResolveAttack` without prompting. Coding alongside each lesson consistently.

## Zone of proximal development
Ready for: goroutines and channels (Go concurrency), or spell slots / resource tracking (D&D). Also: type assertions (`pc.(*PlayerCharacter)`) will come up naturally once the combat loop needs to distinguish PCs from monsters for death saves.
