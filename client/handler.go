package client

import (
	"regexp"

	"github.com/amarnathcjd/gogram/telegram"
)

type Handler func(message *telegram.NewMessage) error
type CallbackHandler func (message *telegram.CallbackQuery) error 

const CMD_PREFIXES = "!.?"
var cCMDS = make(map[string]CallbackHandler)

func RegCallback(command string, handler CallbackHandler) {
	UserBot.AddCallbackHandler(command, handler)
	cCMDS[command] = handler
}

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
