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

func setLastNameHandler(m *tg.NewMessage) error {
	name := m.Args()
	if name == "" {
		return EoR(m, "Please Provide A Name")
	}
	_, err := client.UserBot.AccountUpdateProfile("", name, "")
	if err != nil {
		return EoR(m, err.Error())
	}
	return EoR(m, "Last Name Changed!")
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

func setPfpHandler(m *tg.NewMessage) error {
	if !m.IsReply() {
		return EoR(m, "Please Reply To A Photo")
	}
	f, err := m.Download()
	if err != nil {
		return EoR(m, err.Error())
	}
	inpf, err := m.Client.UploadFile(f)
	if err != nil {
		return EoR(m, err.Error())
	}
	_, err = client.UserBot.PhotosUploadProfilePhoto(inpf, nil, 0)
	if err != nil {
		return EoR(m, err.Error())
	}
	return EoR(m, "Profile Picture Changed!")
}

func init() {
	client.RegCmd("setname", setNameHandler)
	client.RegCmd("setlastname", setLastNameHandler)
	client.RegCmd("setbio", setBioHandler)
	client.RegCmd("setpfp", setPfpHandler)
}
