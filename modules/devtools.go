package modules

import (
	"fmt"
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
	fmt.Println(out, err)
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

func init() {
	client.RegCmd("shell", shellHandler)
}
