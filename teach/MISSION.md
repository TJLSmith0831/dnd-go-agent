# Mission: Go + D&D 5e via Building a DM Bot

## Why
Build a production-quality Discord/Slack Dungeon Master bot in Go as the vehicle for learning two things simultaneously: idiomatic Go (with real systems design — concurrency, caching, scaling) and D&D 5e rules deeply enough to implement them correctly in code. The project is the curriculum.

## Success looks like
- Can write idiomatic Go without falling back on Python patterns (goroutines, interfaces, error handling feel natural)
- Understands Go's concurrency model well enough to design the agent loop and message pipeline confidently
- Knows D&D 5e core rules (combat, ability checks, spellcasting, character progression) well enough to implement them without referencing the rulebook for every decision
- Has made real systems design decisions in the codebase — at least one around caching, one around scaling — and can explain the tradeoffs
- The bot can run a simple one-shot D&D session end-to-end in Discord or Slack

## Constraints
- Max 3 hours per session
- Comes from Python/TS/JS background — frame Go explanations via Python analogues, highlight where Go idioms differ
- Already understands LLM API calls and agent architecture (LangGraph/LangChain) — no need to teach those
- Postgres-comfortable — DB design doesn't need first-principles treatment

## Out of scope (for now)
- Deep LLM/agent architecture (already knows this)
- Advanced D&D content (subclasses, multiclassing, magic item crafting) — learn core rules first
- Frontend/UI beyond Discord and Slack Bolt
- Premature performance optimization before the bot can run a session
