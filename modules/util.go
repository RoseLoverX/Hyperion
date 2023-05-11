package modules

import (
	"main/client"

	tg "github.com/amarnathcjd/gogram/telegram"
)

func EoR(m *tg.NewMessage, msg string) error {
	if client.UserBot.CommanderId() == client.UserBot.SelfId() {
		_, err := m.Edit(msg)
		return err
	} else {
		_, err := m.Reply(msg)
		return err
	}
}

func EorW(m *tg.NewMessage, msg string) (*tg.NewMessage, error) {
	if client.UserBot.CommanderId() == client.UserBot.SelfId() {
		return m.Edit(msg)
	} else {
		return m.Reply(msg)
	}
}
