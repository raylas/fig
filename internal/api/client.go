package api

import (
	"github.com/sashabaranov/go-openai"
)

type Client struct {
	*openai.Client
	model  string
	tokens int
}

func NewClient(apiKey, model string, tokens int) *Client {
	return &Client{
		Client: openai.NewClient(apiKey),
		model:  model,
		tokens: tokens,
	}
}

func (c *Client) Query(message string) error {
	r := buildRequest(c.model, message, c.tokens, true)
	err := c.Stream(r)
	if err != nil {
		return err
	}
	return nil
}

func buildRequest(model, message string, tokens int, stream bool) openai.ChatCompletionRequest {
	return openai.ChatCompletionRequest{
		Model:     model,
		MaxTokens: tokens,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: message,
			},
		},
		Stream: stream,
	}
}
