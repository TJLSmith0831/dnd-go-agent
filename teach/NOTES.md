# Teaching Notes

## User Profile
- Strong Python, TS/JS, Postgres, R background
- Experienced with LangGraph/LangChain and LLM APIs — skip agent architecture basics
- Learning Go and D&D 5e simultaneously via this project
- Sessions max 3 hours

## Bot Connection
- **Every lesson must visibly move the bot forward.** Each lesson should state what the bot can now do that it couldn't before. If a Go concept can't be tied to a concrete bot capability, defer it or cut it.
- Keep `main.go` up to date with each lesson — it is the running definition of what the bot currently is.

## Teaching Preferences
- Frame Go idioms against Python — show the delta, not from zero
- Dual-track: Go concepts + D&D 5e rules, tied to the bot project
- Systems design (caching, concurrency, scaling) is an explicit goal — surface tradeoffs, don't hide them
- **TDD is preferred workflow** — write failing tests first, then implement. Bake red→green into every lesson that introduces new code. Don't leave tests as a post-lesson exercise.
- When test setups use structs, explicitly call out which fields need non-zero values — zero-value bugs (maxHP=0, damageDie=0) caused real test failures and needed teacher intervention.

## Quiz Design
- **All answer options must be the same length** — user noticed the correct answer was always the longest. Don't give away answers through formatting. Aim for equal word count and character count across all four options for every question.

## D&D 5e Experience
- Has played BG3 significantly — knows action economy, combat loop, ability checks at an intuitive level
- Has DMed before (self-assessed: not very good) — understands the narrative/encounter side
- Does NOT need first-principles intro to D&D — can jump straight to formalizing the rules mechanically
- BG3 is faithful 5e, so concepts like Advantage, AC, saving throws, spell slots are already familiar
