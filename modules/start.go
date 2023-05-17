package modules

import (
	"fmt"
	"main/client"
	"strings"
	"time"

	tg "github.com/amarnathcjd/gogram/telegram"
)

func AddUser(m *tg.NewMessage) error {
	if m.SenderID() != client.UserBot.CommanderId() {
		m.Reply("You are not my commander!")
		return nil
	}
	ux := m.Args()
	if ux == "" {
		return EoR(m, "Please provide a user id and chat id seperated by space")
	}
	uxx := strings.Split(ux, " ")
	if len(uxx) != 2 {
		return EoR(m, "Please provide a user id and chat id seperated by space")
	}
	userId := uxx[0]
	chatId := uxx[1]

	userPeer, err := client.UserBot.GetSendablePeer(userId)
	if err != nil {
		m.Reply("Error: " + err.Error())
		return nil
	}
	chatPeer, err := client.UserBot.GetSendablePeer(chatId)
	if err != nil {
		m.Reply("Error: " + err.Error())
		return nil
	}

	if peerChannel, ok := chatPeer.(*tg.InputPeerChannel); ok {
		if peerUser, ok := userPeer.(*tg.InputPeerUser); ok {
			_, err := client.UserBot.ChannelsInviteToChannel(&tg.InputChannelObj{
				ChannelID:  peerChannel.ChannelID,
				AccessHash: peerChannel.AccessHash,
			},
				[]tg.InputUser{
					&tg.InputUserObj{
						UserID:     peerUser.UserID,
						AccessHash: peerUser.AccessHash,
					},
				},
			)

			if err != nil {
				return EoR(m, err.Error())
			}
		} else {
			return EoR(m, "Please provide a valid user id")
		}
	} else {
		return EoR(m, "Please provide a valid chat id")
	}

	return EoR(m, "User Added!")
}

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
	peer, err := client.UserBot.GetSendablePeer(link)
	if err != nil {
		if strings.Contains(link, "t.me/+") {
			link = strings.ReplaceAll(link, "t.me/+", "")
			if strings.HasPrefix(link, "https://") {
				link = strings.ReplaceAll(link, "https://", "")
			}
			if strings.HasPrefix(link, "http://") {
				link = strings.ReplaceAll(link, "http://", "")
			}
			_, err := client.UserBot.MessagesImportChatInvite(link)
			if err != nil {
				return EoR(m, err.Error())
			}
			return EoR(m, "Joined!")
		}
		return EoR(m, err.Error())
	}
	if peerChannel, ok := peer.(*tg.InputPeerChannel); ok {
		_, err := client.UserBot.ChannelsJoinChannel(&tg.InputChannelObj{
			ChannelID:  peerChannel.ChannelID,
			AccessHash: peerChannel.AccessHash,
		})
		if err != nil {
			return EoR(m, err.Error())
		}
	} else {
		return EoR(m, "Please provide a valid channel link")
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
	client.RegCmd("adduser", AddUser)
}
