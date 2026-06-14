# CLAUDE.md - dnd-go-agent

Multi-agent D&D bot in Go. Two modes sharing one orchestrator: **DM mode** (Discord, full fantasy campaigns with Narrator, NPC, and Referee agents and persistent world state) and **Productivity mode** (Slack, work events mapped to quest narrative: completed tickets become "monsters slain," standups become story beats).

This is a learning project. Optimize for understanding idiomatic Go and the agent loop, not speed or cleverness.

**Tradeoff:** These guidelines bias toward caution and learning over raw velocity. For trivial tasks (typos, formatting), use judgment.

## 1. Think Before Coding

**Don't assume. Don't hide confusion. Surface tradeoffs.**

- State assumptions explicitly, especially around Go idioms (interfaces, goroutines, channels, error handling). If unsure how something "should" be done in Go, ask or flag it as a learning moment.
- If multiple approaches exist, present them with tradeoffs rather than picking silently.
- Push back if a simpler approach exists.

## 2. Simplicity First

**Minimum code that solves the problem. Nothing speculative.**

- Build only what the current mode/feature needs. No speculative abstractions for "future modes."
- Don't build a generic plugin system, config DSL, or interface layer until at least two concrete cases need it.
- Shared state (campaign/world DB) should be the simplest schema that works for both modes. Don't over-normalize early.

## 3. Surgical Changes

**Touch only what you must. Clean up only your own mess.**

- Touch only what the current task requires.
- Match existing Go style and package structure; don't reorganize unless asked.
- Flag dead code or TODOs you notice, don't remove unless asked.

## 4. Goal-Driven Execution

**Define success criteria. Loop until verified.**

For each task, state a short plan with verification steps, e.g.:
```
1. Add Referee agent dice-roll tool -> verify: unit test for roll outcomes
2. Wire Slack /quest command -> verify: manual message round-trip in test workspace
```

## Git Commits

Split commits by authorship:

- **User wrote it** (Go source, go.mod, go.sum, main.go): commit with no co-author.
- **Claude wrote it** (lessons `teach/lessons/*.html`, learning records `teach/learning-records/*.md`, CLAUDE.md changes): commit with `Co-Authored-By: Claude Sonnet 4.6 <noreply@anthropic.com>`.

## Security

**Never read `.env` files.** They may contain API keys and secrets. If context about environment variables is needed, ask the user to describe them — never open the file.

## Project Specifics

- **Language:** Go (learning goal: idiomatic concurrency, interfaces, error handling) and DnD 5e rules
- **Modes:** `dm` and `productivity`, selected via config/flag, sharing one orchestrator and agent set
- **Agents:** Narrator, NPC, Referee, same code, different system prompts/tools per mode
- **Frontends:** Discord (DM mode), Slack via Bolt for Go (Productivity mode)
- **State:** Postgres, single schema covering campaigns and quest events
- **LLM:** Anthropic API directly, or via go-agent/LangChainGo if it simplifies the agent loop without hiding the mechanics

---

**These guidelines are working if:** fewer Rust-style or Python-style patterns leak into the Go code, questions about Go idioms come before implementation, and the agent loop stays simple enough that you can explain how a request flows end to end.

## Agent skills

### Issue tracker

Issues live as local markdown files under `.scratch/<feature>/`. See `docs/agents/issue-tracker.md`.

### Triage labels

Default label vocabulary (`needs-triage`, `needs-info`, `ready-for-agent`, `ready-for-human`, `wontfix`). See `docs/agents/triage-labels.md`.

### Domain docs

Single-context layout — one `CONTEXT.md` + `docs/adr/` at the repo root. See `docs/agents/domain.md`.
