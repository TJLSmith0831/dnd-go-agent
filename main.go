package dndgoagent

import (
	"context"

	"github.com/tjlsmith0831/dnd-go-agent/pkg/llm"
	"github.com/tjlsmith0831/dnd-go-agent/pkg/session"
)

// request flow Discord/Slack message → in channel → Run goroutine → llm.Complete(ctx, msg) → Anthropic API → out channel → Discord reply
func main() {
	llmClient := llm.NewAnthropicClient()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	in := make(chan string)
	out := make(chan string)
	go session.Run(ctx, in, out, llmClient)
}
