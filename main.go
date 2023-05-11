package main

import (
	"fmt"
	"main/client"
	_ "main/modules"
)

func main() {
	fmt.Println(client.UserBot.CommanderId())
	client.UserBot.Idle()
}
