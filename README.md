# dnd-go-agent

A multi-agent D&D bot written in Go. Two modes share one orchestrator:

- **DM mode** — Discord, full fantasy campaigns with Narrator, NPC, and Referee agents and persistent world state
- **Productivity mode** — Slack, work events mapped to quest narrative: completed tickets become "monsters slain," standups become story beats

This is also a learning project. See [`teach/`](./teach/) for the curriculum.

## Packages

| Package | Description |
|---|---|
| `pkg/dice` | Dice rolling, ability modifiers, ability checks |
| `pkg/combat` | Combat engine: initiative, attack resolution, death saves, zero-HP handling |

## Planned

- Narrator, NPC, Referee agents (different system prompts/tools per mode)
- Discord frontend (DM mode) and Slack/Bolt frontend (Productivity mode)
- Postgres state: campaigns, world events, quest log
- Anthropic API agent loop

## Running tests

```sh
go test ./...
```

## Module

```
github.com/tjlsmith0831/dnd-go-agent
```
