---
name: postgres-pgx
description: Lesson 13 — pgx/v5, schema design (sessions + events), Exec/QueryRow/Query, rows.Close/Err, Store interface with stub + integration test.
metadata:
  type: feedback
---

## What was covered
- pgx/v5 vs database/sql: pgx is Postgres-native (UUID, JSONB, arrays, batch queries)
- `pgx.Connect(ctx, connStr)` — context-cancellable connection
- `conn.Exec` — INSERT/UPDATE/DELETE, no rows returned
- `conn.QueryRow(...).Scan(&v)` — single row; `pgx.ErrNoRows` sentinel when no match
- `conn.Query` — multiple rows; must `defer rows.Close()` immediately after error check
- `rows.Next()` + `rows.Scan()` loop; `rows.Err()` after loop catches mid-stream errors
- `$1, $2...` positional params (same as psycopg2)
- `Store` interface in `pkg/store` (consumer-owned): `AppendEvent`, `RecentEvents`
- Integration test pattern: skip if `DATABASE_URL` not set (mirrors `pkg/llm`)

## Schema
Two tables: `sessions` (one per guild, `guild_id UNIQUE`) and `events` (one per message, role ∈ player/narrator/referee/npc). Events table IS the campaign journal and the Narrator's long-term memory.

## Bot capability unlocked
Store every player action and narrator response. On reconnect, load last N events and inject as message history → Narrator remembers previous sessions.

## main.go fix
`package dndgoagent` → `package main`. Added stdin loop so `go run .` actually runs the bot. Loads .env via godotenv.

## Zone of proximal development
Ready for: Lesson 14 — system prompts and conversation history. Wire the store into session.Run so the Narrator receives full message history per turn, and give each agent a defining system prompt.
