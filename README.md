# fig

**fig(ment)** | ˈfiɡmənt |
```
a thing that someone believes to be real but that exists only in their imagination: it really was Ross and not a figment of her overheated imagination.
```

A not-very-exciting CLI tool to query or chat with the [OpenAI API](https://platform.openai.com/docs/api-reference).

[![Main](https://github.com/raylas/fig/actions/workflows/main.yaml/badge.svg)](https://github.com/raylas/fig/actions/workflows/main.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/raylas/fig)](https://goreportcard.com/report/github.com/raylas/fig)

## Configuration

- `OPENAI_API_KEY` (required: [docs](https://platform.openai.com/docs/api-reference/authentication))
- `OPENAI_MAX_TOKENS` (default: `500`)

## Usage

```sh
Usage: main [-a APIKEY] [-t TOKENS] [-m MODEL] <command> [<args>]

Options:
  -a APIKEY              OpenAI API key [env: OPENAI_API_KEY]
  -t TOKENS              maximum number of tokens to return [default: 500, env: OPENAI_MAX_TOKENS]
  -m MODEL               model to use [default: gpt-3.5-turbo]
  --help, -h             display this help and exit

Commands:
  query                  query the model
  chat                   chat with the model
```
