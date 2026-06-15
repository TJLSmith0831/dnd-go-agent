---
name: skills-proficiency
description: Lesson 10 — D&D 5e skills (18 skills mapped to abilities), proficiency bonus by level, expertise, passive checks, Referee decision tree.
metadata:
  type: feedback
---

## What was covered
- 18 skills each anchored to one ability score (STR/DEX/INT/WIS/CHA — CON has none)
- Proficiency bonus by character level: +2 at 1–4, +3 at 5–8, +4 at 9–12, +5 at 13–16, +6 at 17–20
- Skill check: d20 + ability modifier + proficiency bonus (if proficient) ≥ DC
- Non-proficient: d20 + ability modifier only
- Expertise (Rogues, Bards): double proficiency bonus on selected skills
- Passive check: 10 + ability modifier (+ proficiency if applicable) — no roll, used for automatic detection
- Passive Perception: the fixed score that determines what characters notice without actively looking
- Referee decision tree: (1) is a check needed? (2) which skill? (3) proficient? (4) what DC?

## TDD exercise
Add `SkillCheck(score, profBonus int, proficient bool, dc int) (CheckResult, error)` and `PassiveCheck(score, profBonus int, proficient bool) int` to `pkg/dice`. Expertise handled by caller passing `profBonus * 2`.

## Bot connection
`pkg/dice` now has everything the Referee agent needs: `AbilityCheck`, `SkillCheck`, `PassiveCheck`, `ResolveAttack`. The Referee's job is mapping free-text player intent → the right function call.

## Zone of proximal development
Ready for: Lesson 11 — Go error handling (`errors.Is`, `errors.As`, `fmt.Errorf("%w")`, sentinel errors). Needed for the Referee agent to surface "no skill match" vs "dice roll failed" distinctly.
