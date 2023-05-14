package client

import (
	"regexp"

	"github.com/amarnathcjd/gogram/telegram"
)

type Handler func(message *telegram.NewMessage) error

const CMD_PREFIXES = "!."

var CMDS = make(map[string]Handler)

func RegCmd(command string, handler Handler) {
	UserBot.AddMessageHandler(regexp.MustCompile(`(?i)^[`+CMD_PREFIXES+`]`+command), handler, &telegram.Filters{
		Users: UserBot.CommanderId(),
	})
	CMDS[command] = handler
}
