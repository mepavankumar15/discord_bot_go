package main

import (
	"fmt"

	"github.com/mepavankumar15/discord_bot_go/bot"
	"github.com/mepavankumar15/discord_bot_go/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}
