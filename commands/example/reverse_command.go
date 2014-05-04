package example

import (
	"github.com/fabioxgn/go-bot/cmd"
)

// From stackoverflow: http://stackoverflow.com/a/10030772
func reverse(command *cmd.Cmd) (msg string, err error) {
	runes := []rune(command.FullArg)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	msg = string(runes)
	return
}

func init() {
	cmd.RegisterCommand(&cmd.CustomCommand{
		Cmd:         "reverse",
		CmdFunc:     reverse,
		Description: "Reverses the whole string",
		Usage:       "all your base are belong to us",
	})
}