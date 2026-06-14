package session

import (
	"context"
	"testing"
)

type stubLLM struct{}

func (s stubLLM) Complete(_ context.Context, _ string) (string, error) {
	return "You swing and miss.", nil
}

func TestRun(t *testing.T) {
	in := make(chan string)
	out := make(chan string)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go Run(ctx, in, out, stubLLM{})

	in <- "attack goblin"
	got := <-out
	if got != "You swing and miss." {
		t.Errorf("got %q, want %q", got, "You swing and miss.")
	}
}
