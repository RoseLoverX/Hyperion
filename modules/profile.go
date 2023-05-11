package modules

import (
	"main/client"

	tg "github.com/amarnathcjd/gogram/telegram"
)

func setNameHandler(m *tg.NewMessage) error {
	name := m.Args()
	if name == "" {
		return EoR(m, "Please Provide A Name")
	}
	_, err := client.UserBot.AccountUpdateProfile(name, "", "")
	if err != nil {
		return EoR(m, err.Error())
	}
	return EoR(m, "Name Changed!")
}

func setBioHandler(m *tg.NewMessage) error {
	bio := m.Args()
	if bio == "" {
		return EoR(m, "Please Provide A Bio")
	}
	_, err := client.UserBot.AccountUpdateProfile("", bio, "")
	if err != nil {
		return EoR(m, err.Error())
	}
	return EoR(m, "Bio Changed!")
}

func init() {
	client.RegCmd("setname", setNameHandler)
	client.RegCmd("setbio", setBioHandler)
}
