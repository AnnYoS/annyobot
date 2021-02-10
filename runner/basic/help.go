package basic

import (
	"annyobot/command/model"
	"github.com/bwmarrin/discordgo"
)

func GetStart(session *discordgo.Session, message *discordgo.Message, param string, options []model.Option) {
	msg := "```You want to know basics about me ? Alright !\n" +
		"All my commands begin with '!', so when you doing this, I going to talk in the channel.\n" +
		"If you want to know more about me, try !help command, it's free.```"
	_, _ = session.ChannelMessageSend(message.ChannelID, msg)
}
