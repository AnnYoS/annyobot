package basic

import (
	"annyobot/command"
	"github.com/bwmarrin/discordgo"
)

func GetStart(session *discordgo.Session, message *discordgo.Message, param string, options []command.Option) {
	msg := "```You want to know basics about me ? Alright !\n" +
		"All my commands begin with '!', so when you doing this, I going to talk in the channel.\n" +
		"If you want to know more about me, try !help command, it's free.```"
	_, _ = session.ChannelMessageSend(message.ChannelID, msg)
}

func Help(session *discordgo.Session, message *discordgo.Message, param string, options []command.Option) {

	msg := "```\nPrototype of a command: \n" +
		"!command <Parameter> <Option1> <Option2>...\n\n"
	if param != "" {
		c := command.FindCommand("!" + param)
		if c == nil {
			_, _ = session.ChannelMessageSend(message.ChannelID, "Cannot giving help for this command")
		}
		msg += c.Command + ": " + c.Desc + "\nParameter: " + c.Param + "\nOption(s):\n"
		for _, o := range c.Options {
			msg += "   " + o.Opt + ": " + o.Desc + "\n"
		}
	} else {
		for _, c := range command.Association {
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
