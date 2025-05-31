package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/mepavankumar15/discord_bot_go/config"
)

var BotID string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	BotID = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running !")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, `hello guys your talking to the hero of server 
		avyukth an discord bot`)
	}
	if m.Content == "how are you" {
		_, _ = s.ChannelMessageSend(m.ChannelID, `the hero is alright,  how may i help you --avyukth an discord bot`)
	}
	if m.Content == "tell me a joke" {
		_, _ = s.ChannelMessageSend(m.ChannelID, `your so smart , i think you got the joke 😉`)
	}
}
