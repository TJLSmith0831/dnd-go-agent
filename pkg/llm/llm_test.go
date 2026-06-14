package llm

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

// Compile-time check: *AnthropicClient satisfies the LLMClient shape
// that session.Run expects. This costs nothing at runtime — assigning nil
// to a typed interface variable is enough for the compiler to verify the
// method set. If AnthropicClient's Complete signature ever drifts, this
// line fails to compile before any test runs.
var _ interface {
	Complete(ctx context.Context, prompt string) (string, error)
} = (*AnthropicClient)(nil)

// TestAnthropicClientIntegration hits the real API.
// Skipped automatically when ANTHROPIC_API_KEY is not set, so go test ./...
// stays green in CI and offline. Run manually to verify end-to-end wiring:
//
//	ANTHROPIC_API_KEY=sk-... go test ./pkg/llm/ -run Integration -v
func TestAnthropicClientIntegration(t *testing.T) {
	godotenv.Load("../../.env")
	if os.Getenv("ANTHROPIC_API_KEY") == "" {
		t.Skip("ANTHROPIC_API_KEY not set")
	}

	client := NewAnthropicClient()
	ctx := context.Background()

	response, err := client.Complete(ctx, "Reply with exactly three words.")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if response == "" {
		t.Error("expected non-empty response, got empty string")
	}
	t.Logf("response: %q", response)
}
