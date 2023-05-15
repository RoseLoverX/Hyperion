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
		Func: filterUserAuthorized,
	})
	CMDS[command] = handler
}

func filterUserAuthorized(message *telegram.NewMessage) bool {
	senderId := message.SenderID()
	if senderId == UserBot.CommanderId() {
		return true
	}
	if IsCachedSudoer(senderId) {
		return true
	}
	return false
}
