package api

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (c *Client) Chat() error {
	fmt.Println("Type `quit` to exit.")
	for {
		s := prompt()
		if s == "quit" {
			break
		}
		err := c.Query(s)
		if err != nil {
			return err
		}
		fmt.Printf("\n")
	}
	return nil
}

func prompt() string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, "> ")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}
