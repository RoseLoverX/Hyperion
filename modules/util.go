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
		_, err := m.Reply(msg, tg.SendOptions{ParseMode: "HTML"})
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

func PrettyPrint(m *tg.NewMessage, msg string) error {
	if client.UserBot.CommanderId() == client.UserBot.SelfId() {
		_, err := m.Edit("<pre>" + msg + "</pre>")
		return err
	} else {
		_, err := m.Reply("<pre>"+msg+"</pre>", tg.SendOptions{ParseMode: "HTML"})
		return err
	}
}
