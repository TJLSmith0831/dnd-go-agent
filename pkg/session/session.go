package session

import "context"

/*
In Python, you reach for asyncio or threads when you need to do two things at once.
Go has a different answer: goroutines and channels. These aren't a library bolt-on,
and they're built into the language. Understanding them is the biggest conceptual jump
from Python to Go, and it's the foundation of everything the bot does concurrently:
handling multiple Discord guilds, routing agent messages, processing dice rolls,
and processing a narrative response.
*/

/*
## Goroutines
A goroutine is a function running concurrently with the rest of your program. You launch
one with the go keyword: `go someFunction()`. That's it. The function starts immediately in the background.
Your current code keeps running.

### How goroutines differ from threads
OS threads are expensive: each one reserves ~ 1MB of stack space and requires a kernel context switch.
Goroutines start with a ~2KB stack that grows as needed and are multiplexed onto a small pool of
OS threads by the Go runtime scheduler. You can run tens of thousands of goroutines in a single process
without issue — the bot can maintain one goroutine per active campaign session with no sweat.

### Goroutine lifecycle
A goroutine exits when its function returns. There's no handle you can join or cancel directly —
you signal it through a channel. This matters: a goroutine that's blocked waiting on a channel with
no sender will leak and never exit. The Go race detector (got test -race) can help catch these leaks.
*/

/*
## Channels
Goroutines communicate through channels. A channel is a typed pipe: you send
a value in from one end and receive it out the other. The Go mantra is:
"Don't communicate by sharing memory; share memory by communicating." ~ Rob Pike
^^ Instead of a shared variable protected by a mutex, pass values through a channel.
   One goroutine owns the value; it hands it off via the channel. No lock needed.

### Creating and using a channel
```go
	// make(chan T) creates an unbuffered channel of type T
	ch := make(chan string)

	// send: ch <- value   (blocks until a receiver is ready)
	ch <- "hello"

	// receive: value := <-ch   (blocks until a sender sends)
	msg := <-ch
```

### Unbuffered channels block on both ends
The sender waits until a receiver is ready; the receiver waits until a sender sends.
This synchronization is the whole point — it's a rendezvous, not a queue.

### Directional channels
You can restrict a channel to send-only or receive-only:

```go
// send-only channel
ch := make(chan<- string)

// receive-only channel
ch := make(<-chan string)
```

This is a compile-time guarantee: you can't send on a receive-only channel or
receive on a send-only channel.
A bidirectional chan string converts implicitly to either directional form — you make it
at the call site and pass it in restricted.

### Ranging over a channel
A for range loop on a channel reads values until the channel is closed. This is the standard
agent loop pattern:

```go
for msg := range in {
    // process each message as it arrives
    out <- process(msg)
}
// loop exits only when in is closed
```
The loop blocks between messages — the goroutine is suspended until the next value arrives.
Closing the channel from the sender side signals "no more values." Closing from the receiver
side is a panic — only the sender closes.
*/

/*
## Select Statements
`select` is like a switch for channel operations. It blocks until one of its cases is ready,
then executes that case. If multiple are ready, it picks one at random.

```go
select {
case msg := <-playerMsg:
    handlePlayerMessage(msg)
case <-done:
    return // shutdown signal received
}
```

The done channel pattern is the standard way to cancel a goroutine: the parent closes a
`chan struct{}` (zero-size type, costs nothing), and the goroutine's select picks it up and returns.

*/

/*
How this all maps to the bot:
- One goroutine per Discord server (guild) / Slack workspace that's running a game.
- It loops over an incoming message channel with `for range` and a `done` channel in select.
- The orchestrator sends player messages in; the agent loop reads them out and sends
  LLM responses back through an outbound channel. Campaign sessions start and stop by opening and closing channels.
*/

// in is <-chan string — receive-only, Run can't accidentally send to it
// out is chan<- string — send-only, Run can't accidentally read from it
// done is <-chan struct{} — receive-only; struct{} costs zero bytes
// func Run(in <-chan string, out chan<- string, done <-chan struct{}) {
// 	for {
// 		select {
// 		case msg := <-in:
// 			out <- msg
// 		case <-done:
// 			return
// 		}
// 	}
// }

/*
## context.Context

`context.Context` is a standard way to pass cancellation signals, deadlines, and key-value pairs
between goroutines. It's always the first parameter in blocking operations.

You'll see it as the first parameter of almost every function in Go that does I/O:

```go
func Run(ctx context.Context, ...) {}
func QueryDB(ctx context.Context, ...) {}
func Complete(ctx context.Context, ...) {}
```

First parameter, always. It's a Go convention, not a compiler requirement — but the entire
ecosystem follows it so consistently that violating it feels wrong immediately.

### Creating a context
```go
// Background: the root context. Used in main or top-level tests.
ctx := context.Background()

// WithCancel: derive a child context with a cancel function.
ctx, cancel := context.WithCancel(context.Background())
defer cancel() // always call cancel — releases resources
```

`context.WithCancel` returns two values: a new context derived from the parent, and a cancel
function. Calling `cancel()` closes `ctx.Done()` — a `<-chan struct{}`. Sound familiar? It's
the same done-channel pattern you wrote in Lesson 8, standardised.

> Your done channel, upgraded:
`done <-chan struct{}` is the primitive. context.Context is the abstraction built on top of it. Both
work BUT context.Context composes better. If `Run` calls another function that takes a context, cancellation
propagates automatically. With a raw `done` channel, you'd have to pass it down manually.

### Listening for cancellation

```go
select {
case msg := <-in:
    // handle message
case <-ctx.Done():
    return // context cancelled — clean exit
}
```

### Passing context into blocking calls
When you call the Anthropic API, the HTTP request might take seconds. Passing `ctx` into the SDK
call means a cancelled context aborts the in-flight request immediately. You don't leak a goroutine
or waste API credits on a response nobody's waiting for.

```go
response, err := llm.Complete(ctx, msg) // ctx cancels the HTTP call if the session ends
```
*/

type LLMClient interface {
	Complete(ctx context.Context, prompt string) (string, error)
}

func Run(ctx context.Context, in <-chan string, out chan<- string, llm LLMClient) {
	for {
		select {
		case msg := <-in:
			response, err := llm.Complete(ctx, msg)
			if err != nil {
				continue
			}
			out <- response
		case <-ctx.Done():
			return
		}
	}
}
