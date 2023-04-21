package api

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/sashabaranov/go-openai"
)

func (c *Client) Stream(req openai.ChatCompletionRequest) error {
	ctx := context.Background()

	stream, err := c.CreateChatCompletionStream(ctx, req)
	if err != nil {
		return err
	}
	defer stream.Close()

	for {
		resp, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			return nil
		}
		if err != nil {
			return err
		}

		fmt.Printf(resp.Choices[0].Delta.Content)
	}
}
