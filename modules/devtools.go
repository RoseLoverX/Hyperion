package modules

import (
	"main/client"
	"os/exec"

	tg "github.com/amarnathcjd/gogram/telegram"
)

func shellHandler(m *tg.NewMessage) error {
	cmd := m.Args()
	if cmd == "" {
		return EoR(m, "Please Provide A Command")
	}
	out, err := shell(cmd)
	if err != nil {
		return EoR(m, err.Error())
	}
	return EoR(m, "<code>"+out+"</code>")
}

func shell(cmd string) (string, error) {
	command := exec.Command("bash", "-c", cmd)
	out, err := command.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func getPreview(m *tg.NewMessage) error {
	url := m.Args()
	req, err := client.UserBot.MessagesGetWebPage(url, 0)
	if err != nil {
		m.Reply("Error: "+ err.Error())
		return nil
	}
	b, _ := json.Marshal(req)
	m.Reply("<code>" + string(b) + "</code>")
	return nil
}

func init() {
	client.RegCmd("shell", shellHandler)
	client.RegCmd("preview", getPreview)
}
