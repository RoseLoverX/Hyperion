package modules

import (
	"fmt"
	"main/client"
	"time"

	tg "github.com/amarnathcjd/gogram/telegram"
)

func startHandler(m *tg.NewMessage) error {
	return EoR(m, "Hyperion Userbot Is Running")
}

func pingHandler(m *tg.NewMessage) error {
	initTime := time.Now()
	msg, err := EorW(m, "Pong!")
	if err != nil {
		return err
	}
	latency := time.Since(initTime).Milliseconds()
	_, err = msg.Edit(fmt.Sprintf("<b>Pong!</b> %dms <b>||</b> <b>Uptime:</b> <spoiler>%s</spoiler>", latency, time.Since(time.Unix(client.StartTime, 0)).String()))
	return err
}

func joinChannelHandler(m *tg.NewMessage) error {
	link := m.Args()
	if link == "" {
		return EoR(m, "Please Provide A Link")
	}
	err := client.UserBot.JoinChannel(link)
	if err != nil {
		return EoR(m, err.Error())
	}
	return EoR(m, "Joined!")
}

func leaveChannelHandler(m *tg.NewMessage) error {
	link := m.Args()
	if link == "" {
		return EoR(m, "Please Provide A Link")
	}
	err := client.UserBot.LeaveChannel(link)
	if err != nil {
		return EoR(m, err.Error())
	}
	return EoR(m, "Left!")
}

func init() {
	client.RegCmd("start", startHandler)
	client.RegCmd("ping", pingHandler)
	client.RegCmd("join", joinChannelHandler)
	client.RegCmd("leave", leaveChannelHandler)
}
