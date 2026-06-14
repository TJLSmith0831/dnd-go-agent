package llm

import (
	"context"
	"fmt"

	anthropic "github.com/anthropics/anthropic-sdk-go"
)

type AnthropicClient struct {
	client anthropic.Client // NewClient returns a value, not a pointer
}

func NewAnthropicClient() *AnthropicClient {
	// NewClient reads ANTHROPIC_API_KEY from the environment automatically.
	return &AnthropicClient{client: anthropic.NewClient()}
}

func (a *AnthropicClient) Complete(ctx context.Context, prompt string) (string, error) {
	msg, err := a.client.Messages.New(ctx, anthropic.MessageNewParams{
		Model:     anthropic.ModelClaudeSonnet4_6,
		MaxTokens: 1024,
		Messages: []anthropic.MessageParam{
			anthropic.NewUserMessage(anthropic.NewTextBlock(prompt)),
		},
	})
	if err != nil {
		return "", err
	}
	if len(msg.Content) == 0 {
		return "", fmt.Errorf("empty response from API")
	}
	return msg.Content[0].AsText().Text, nil
}
