package main

import (
	"fmt"

	"github.com/alexflint/go-arg"
	"github.com/raylas/fig/internal/api"
)

type QueryCmd struct {
	Message string `arg:"positional,required" help:"query, or message, to send to the model"`
}

type ChatCmd struct{}

type args struct {
	ApiKey string    `arg:"-a,--,env:OPENAI_API_KEY" help:"OpenAI API key"`
	Tokens int       `arg:"-t,--,env:OPENAI_MAX_TOKENS" default:"500" help:"maximum number of tokens to return"`
	Model  string    `arg:"-m,--" default:"gpt-3.5-turbo" help:"model to use"`
	Query  *QueryCmd `arg:"subcommand:query" help:"query the model"`
	Chat   *ChatCmd  `arg:"subcommand:chat" help:"chat with the model"`
}

// Set by goreleaser.
var version = "0.0.0"
var build = "dev"

func (args) Version() string {
	return fmt.Sprintf("fig %s-%s", version, build)
}

func main() {
	var args args
	arg.MustParse(&args)
	c := api.NewClient(args.ApiKey, args.Model, args.Tokens)

	switch {
	case args.Query != nil:
		err := c.Query(args.Query.Message)
		if err != nil {
			panic(err)
		}
	case args.Chat != nil:
		err := c.Chat()
		if err != nil {
			panic(err)
		}
	}
}
