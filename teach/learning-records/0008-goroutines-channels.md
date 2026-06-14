---
name: goroutines-channels
description: Lesson 08 — goroutines, channels (directional, range, select), done-channel cancellation, pkg/session TDD exercise.
metadata:
  type: feedback
---

## What was covered
- Goroutines: `go fn()` launches concurrently; runtime multiplexes onto OS threads; cheap (~2KB stack vs ~1MB thread)
- Channels: `make(chan T)`, send `ch <- v` blocks until receiver ready, receive `v := <-ch` blocks until sender sends
- Unbuffered = rendezvous, not a queue
- Directional channels: `<-chan T` (receive-only), `chan<- T` (send-only) — enforce intent at compile time
- `for msg := range ch` — blocks between messages, exits only when channel is closed by sender
- `select` — like switch for channel ops; blocks until one case is ready; random choice if multiple ready
- Done-channel pattern: `chan struct{}` (zero-size); parent closes it, goroutine's `select` picks it up and returns
- Go mantra: don't communicate by sharing memory; share memory by communicating
- Data race: concurrent map writes are undefined behavior — race detector (`go test -race`) catches them

## TDD exercise
Build `pkg/session`: `Run(in <-chan string, out chan<- string, done <-chan struct{})` — the skeleton of the per-guild bot loop. Test first: send a message, assert it echoes back, close done.

## Bot connection
One goroutine per Discord guild. Incoming player messages flow through `in`; agent responses flow back through `out`; `done` closes the session. This is the shape of the entire orchestration layer.

## Zone of proximal development
Ready for: Lesson 09 — wire the session goroutine to a real Anthropic API call; introduce the agent loop with actual LLM responses.
