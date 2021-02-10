package basic

import (
	"annyobot/command"
	"github.com/bwmarrin/discordgo"
)

func Hello(session *discordgo.Session, message *discordgo.Message, param string, options []command.Option) {
	if len(options) > 1 {
		_, _ = session.ChannelMessageSend(message.ChannelID, "I can't know what you want")

	} else {
		if options == nil {
			_, _ = session.ChannelMessageSend(message.ChannelID, "Hello !")
		} else if options[0].Opt == "-a" {
			_, _ = session.ChannelMessageSend(message.ChannelID, "Hello @everyone")
		} else if options[0].Opt == "-m" {
			_, _ = session.ChannelMessageSend(message.ChannelID, "Hello "+message.Author.Mention())
		}
	}
}
