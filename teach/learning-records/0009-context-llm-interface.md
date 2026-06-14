---
name: context-llm-interface
description: Lesson 09 — context.Context for cancellation, LLMClient consumer-owned interface, Anthropic SDK wiring, pkg/llm.
metadata:
  type: feedback
---

## What was covered
- `context.Context`: cancellation signal + optional deadline + key-value pairs; passed as first param by universal convention
- `context.Background()` — root context used in main/tests
- `context.WithCancel(parent)` — returns child context + cancel func; `defer cancel()` always
- `ctx.Done()` is a `<-chan struct{}`; replaces the done-channel pattern with a composable abstraction
- Passing ctx into blocking calls (SDK, DB) — cancellation aborts in-flight HTTP requests, preventing goroutine leaks
- `LLMClient` interface defined in `pkg/session` (consumer-owned), not in `pkg/llm` — same pattern as `Combatant`
- `stubLLM` in test satisfies the interface implicitly; no import of the Anthropic SDK needed in tests
- `pkg/llm.AnthropicClient` wraps the SDK; `NewClient()` reads `ANTHROPIC_API_KEY` from env automatically
- Never commit API keys; never store context in a struct — pass it through function params

## TDD exercise
Rewrote `session_test.go` with `stubLLM`, updated `Run` signature to `(ctx context.Context, in, out, llm LLMClient)`, replaced `done` channel with `ctx.Done()` in select.

## Bot connection
The agent loop is now real: player message → in channel → Run goroutine → llm.Complete(ctx, msg) → Anthropic API → out channel → Discord reply. Discord/Slack frontend is what wires the channels to actual messages (Phase 4).

## Zone of proximal development
Ready for: Lesson 10 — D&D skills & proficiency (the last rules layer before the Referee agent can adjudicate any check).
