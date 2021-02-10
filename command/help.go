package command

import (
	"annyobot/command/model"

	"github.com/bwmarrin/discordgo"
)

func Help(session *discordgo.Session, message *discordgo.Message, param string, options []model.Option) {

	msg := "```\nPrototype d'une commande: \n" +
		"!commande <Paramètre> <Option1> <Option2>...\n\n"
	if param != "" {
		c := FindCommand("!" + param)
		if c == nil {
			_, _ = session.ChannelMessageSend(message.ChannelID, "Je ne peux t'aider sur cette commande, sûrement que je n'ai pas cette commande")
		}
		msg += c.Command + ": " + c.Desc + "\nParamètre: " + c.Param + "\nOption(s):\n"
		for _, o := range c.Options {
			msg += "   " + o.Opt + ": " + o.Desc + "\n"
		}
	} else {
		for _, c := range model.Association {
			msg += c.Command + ": " + c.Desc + "\nParamètre: " + c.Param + "\nOption(s):\n"
			for _, o := range c.Options {
				msg += "   " + o.Opt + ": " + o.Desc + "\n"
			}
			msg += "------------------------------------------------\n"
		}
	}
	msg += "```"
	_, _ = session.ChannelMessageSend(message.ChannelID, msg)
}
