package modules

import (
	"fmt"
	"main/client"
	"os/exec"
	"strings"
	"bytes"

	tg "github.com/amarnathcjd/gogram/telegram"
)

func shellHandler(m *tg.NewMessage) error {
	cmd := m.Args()
	if cmd == "" {
		return EoR(m, "Please provide a command")
	}
	out, err := shell(cmd)
	if err != nil && out == "" {
		return EoR(m, "<code>"+err.Error()+"</code>")
	}
	return EoR(m, "<code>"+strings.TrimSpace(strings.TrimPrefix(strings.TrimSuffix(out, "\n"), "\n"))+"</code>")
}

func shell(cmd string) (string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	proc := exec.Command("sh", "-c", args)
	proc.Stdout = &stdout
	proc.Stderr = &stderr
	err := proc.Run()
	return stdout.String() + stderr.String(), err
}

func getPreview(m *tg.NewMessage) error {
	url := m.Args()
	req, err := client.UserBot.MessagesGetWebPage(url, 0)
	if err != nil {
		m.Reply("Error: " + err.Error())
		return nil
	}
	switch req := req.(type) {
	case *tg.WebPageObj:
		caption := req.Title + "\n" + req.Description + "\n<b>Embed Url:</b> " + req.EmbedURL
		if req.Photo != nil {
			m.RespondMedia(req.Photo, tg.MediaOptions{Caption: caption})
		} else if req.Document != nil {
			m.RespondMedia(req.Document, tg.MediaOptions{Caption: caption})
		} else {
			m.Respond(caption)
                }
	case *tg.WebPageEmpty:
		m.Reply("No Preview Found")
	}
	return nil
}

func AddSudoHandler(m *tg.NewMessage) error {
	if m.SenderID() != client.UserBot.CommanderId() {
		m.Reply("You are not my commander!")
		return nil
	}
	var userId int64 = 0
	if m.IsReply() {
		r, err := m.GetReplyMessage()
		if err != nil {
			m.Reply("Error: " + err.Error())
			return nil
		}
		userId = r.SenderID()
	} else {
		userArgs := m.Args()
		if userArgs == "" {
			m.Reply("Please provide a user id")
			return nil
		}
		userPeer, err := client.UserBot.GetSendablePeer(userArgs)
		if err != nil {
			m.Reply("Error: " + err.Error())
			return nil
		}
		userId = client.UserBot.GetPeerID(userPeer)
	}
	if userId == 0 {
		m.Reply("Please provide a valid user id")
		return nil
	}
	if client.IsCachedSudoer(int64(userId)) {
		m.Reply("User already sudoer")
		return nil
	}
	go client.AddSudoer(int64(userId))
	if _, err := m.Reply("User added as sudoer"); err != nil {
		return err
	}
	return nil
}

func RemoveSudoHandler(m *tg.NewMessage) error {
	if m.SenderID() != client.UserBot.CommanderId() {
		m.Reply("You are not my commander!")
		return nil
	}
	var userId int64 = 0
	if m.IsReply() {
		r, err := m.GetReplyMessage()
		if err != nil {
			m.Reply("Error: " + err.Error())
			return nil
		}
		userId = r.SenderID()
	} else {
		userArgs := m.Args()
		if userArgs == "" {
			m.Reply("Please provide a user id")
			return nil
		}
		userPeer, err := client.UserBot.GetSendablePeer(userArgs)
		if err != nil {
			m.Reply("Error: " + err.Error())
			return nil
		}
		userId = client.UserBot.GetPeerID(userPeer)
	}
	if userId == 0 {
		m.Reply("Please provide a valid user id")
		return nil
	}
	if !client.IsCachedSudoer(int64(userId)) {
		m.Reply("User not sudoer")
		return nil
	}
	go client.RemoveSudoer(int64(userId))
	if _, err := m.Reply("User removed as sudoer"); err != nil {
		return err
	}
	return nil
}

func ListSudoHandler(m *tg.NewMessage) error {
	sudoers := client.CACHED_SUDOERS
	if len(sudoers) == 0 {
		m.Reply("No sudoers found")
		return nil
	}
	var msg string = "<b>My Sudoers:</b>\n"
	for _, sudoer := range sudoers {
		msg += fmt.Sprintf("<b>-> <a href=\"tg://user?id=%d\">%d</a></b>\n", sudoer, sudoer)
	}
	if _, err := m.Reply(msg); err != nil {
		return err
	}
	return nil
}

func init() {
	client.RegCmd("shell", shellHandler)
	client.RegCmd("preview", getPreview)

	client.RegCmd("addsudo", AddSudoHandler)
	client.RegCmd("removesudo", RemoveSudoHandler)
	client.RegCmd("listsudo", ListSudoHandler)
}
