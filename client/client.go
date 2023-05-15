package client

import (
	"errors"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/amarnathcjd/gogram/telegram"
	env "github.com/joho/godotenv"
)

const (
	startLogMessage = "Starting Hyperion Userbot..."
)

type Userbot struct {
	*telegram.Client
	commanderId int64
	selfId      int64
	sudoers     []int64
}

var UserBot *Userbot
var StartTime int64 = time.Now().Unix()

func init() {
	var err error
	UserBot, err = InitiallizeUserbot()
	if err != nil {
		panic(err)
	}
}

func NewUserbot(stringSession string) (*Userbot, error) {
	client, err := telegram.NewClient(telegram.ClientConfig{
		StringSession: stringSession,
		MemorySession: true,
		AppHash:       "not_app_hash",
	})
	if err != nil {
		return nil, err
	}
	return &Userbot{
		Client: client,
	}, nil
}

func (u *Userbot) SetCommanderId(id int64) {
	u.commanderId = id
}

func (u *Userbot) CommanderId() int64 {
	return u.commanderId
}

func (u *Userbot) SelfId() int64 {
	return u.selfId
}

func (u *Userbot) Start() error {
	log.Println("Hyperion - Info -", startLogMessage)
	if err := u.Connect(); err != nil {
		return err
	}
	return nil
}

func InitiallizeUserbot() (*Userbot, error) {
	env.Load()
	if stringSession, ok := os.LookupEnv("STRING_SESSION"); ok {
		userbot, err := NewUserbot(stringSession)
		if err != nil {
			return nil, err
		}
		if err := userbot.Start(); err != nil {
			return nil, err
		}
		user, err := userbot.GetMe()
		if err != nil {
			return nil, err
		}
		userbot.selfId = user.ID
		if commanderId, ok := os.LookupEnv("COMMANDER_ID"); ok {
			commanderIdInt64, err := strconv.ParseInt(commanderId, 10, 64)
			if err != nil {
				return nil, errors.New("COMMANDER_ID must be an integer")
			}
			userbot.SetCommanderId(commanderIdInt64)
		} else {
			userbot.SetCommanderId(user.ID)
		}
		return userbot, nil
	}
	return nil, errors.New("STRING_SESSION not found in .env file or environment variables")
}
