package modules

import (
	"fmt"
	"main/client"
	"os/exec"
	"strings"

	tg "github.com/amarnathcjd/gogram/telegram"
)

func shellHandler(m *tg.NewMessage) error {
	cmd := m.Args()
	if cmd == "" {
		return EoR(m, "Please Provide A Command")
	}
	out, err := shell(cmd)
	if err != nil {
		return EoR(m, "<code>"+err.Error()+"</code>")
	}
	return EoR(m, "<code>"+strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(out, "\t", ""), "\r", ""), "\n", "")+"</code>")
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
		m.Reply("Error: " + err.Error())
		return nil
	}
	switch req := req.(type) {
	case *tg.WebPageObj:
		caption := req.Title + "\n" + req.Description + "\n<b>Embed Url:</b> " + req.EmbedURL
		if req.Type == "photo" && req.Photo != nil {
			m.RespondMedia(req.Photo, tg.MediaOptions{Caption: caption})
		} else if req.Document != nil {
			m.RespondMedia(req.Document, tg.MediaOptions{Caption: caption})
		}
	case *tg.WebPageEmpty:
		m.Reply("No Preview Found")
	}
	return nil
}

func AddSudoHandler(m *tg.NewMessage) error {
	if m.SenderID() != client.UserBot.CommanderId() {
		m.Reply("You Are Not My Commander!")
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
			m.Reply("Please Provide A User Id")
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
		m.Reply("Please Provide A User Id")
		return nil
	}
	if client.IsCachedSudoer(int64(userId)) {
		m.Reply("User Already Sudoer")
		return nil
	}
	client.AddSudoer(int64(userId))
	if _, err := m.Reply("User Added As Sudoer"); err != nil {
		return err
	}
	return nil
}

func RemoveSudoHandler(m *tg.NewMessage) error {
	if m.SenderID() != client.UserBot.CommanderId() {
		m.Reply("You Are Not My Commander!")
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
			m.Reply("Please Provide A User Id")
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
		m.Reply("Please Provide A User Id")
		return nil
	}
	if !client.IsCachedSudoer(int64(userId)) {
		m.Reply("User Not Sudoer")
		return nil
	}
	client.RemoveSudoer(int64(userId))
	if _, err := m.Reply("User Removed As Sudoer"); err != nil {
		return err
	}
	return nil
}

func ListSudoHandler(m *tg.NewMessage) error {
	sudoers := client.GetSudoers()
	if len(sudoers) == 0 {
		m.Reply("No Sudoers Found")
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
