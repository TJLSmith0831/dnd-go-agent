---
name: spellcasting
description: Lesson 07 — D&D 5e spellcasting system (slots, resolution paths, concentration, bonus action rule).
metadata:
  type: feedback
---

## What was covered
- Cantrips (0th-level, unlimited) vs leveled spells (consume a slot)
- Spell level ≠ character level; upcasting uses a higher slot for bonus effect
- Full caster slot table through level 9 (1st–5th level slots)
- Two resolution paths: spell attack roll (d20 + spell mod + prof vs AC) and saving throw (d20 + save mod vs DC)
- Spell Save DC formula: 8 + proficiency + spellcasting modifier
- Spellcasting ability by class: INT (Wizard), WIS (Cleric, Druid), CHA (Sorcerer, Bard, Warlock)
- Concentration: one spell at a time, second concentration spell ends first automatically (no save)
- Concentration save on damage: DC = max(10, damage/2), CON save
- Bonus action spell rule: if you cast a spell as a bonus action, your action spell must be a cantrip

## Bot implementation requirements
- Track per caster: spellcasting ability score, proficiency bonus, current/max slots[1–5], active concentration spell
- Trigger concentration save whenever a concentrating PC takes damage
- Validate bonus action spell constraint in action layer

## Zone of proximal development
Ready for: Lesson 08 — goroutines & channels (concurrency model for concurrent sessions and the agent message loop)
