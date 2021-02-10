package command

import (
	"annyobot/command/model"
	"github.com/bwmarrin/discordgo"
)

func Help(session *discordgo.Session, message *discordgo.Message, param string, options []model.Option) {

	msg := "```\nPrototype of a command: \n" +
		"!command <Parameter> <Option1> <Option2>...\n\n"
	if param != "" {
		c := FindCommand("!" + param)
		if c == nil {
			_, _ = session.ChannelMessageSend(message.ChannelID, "Cannot giving help for this command")
		}
		msg += c.Command + ": " + c.Desc + "\nParameter: " + c.Param + "\nOption(s):\n"
		for _, o := range c.Options {
			msg += "   " + o.Opt + ": " + o.Desc + "\n"
		}
	} else {
		for _, c := range model.Association {
			msg += c.Command + ": " + c.Desc + "\nParameter: " + c.Param + "\nOption(s):\n"
			for _, o := range c.Options {
				msg += "   " + o.Opt + ": " + o.Desc + "\n"
			}
			msg += "------------------------------------------------\n"
		}
	}
	msg += "```"
	_, _ = session.ChannelMessageSend(message.ChannelID, msg)
}
