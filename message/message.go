package message

import (
	"annyobot/command"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func ReceiveMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	msg := m.Content
	if !strings.HasPrefix(msg, "!") {
		return
	} else {
		given, err := TransformToCommand(msg)
		if err != nil {
			_, _ = s.ChannelMessageSend(m.ChannelID, "Fait attention à la manière dont tu me donne les options")
			return
		}
		c := command.FindCommand(given.Command)
		if c == nil {
			_, _ = s.ChannelMessageSend(m.ChannelID, "Je ne connais pas la commande, utilise !help pour t'aider :smile:")
			return
		}

		correct := VerifyOptions(given.Options, c.Options)
		if !correct {
			_, _ = s.ChannelMessageSend(m.ChannelID, "Les options que tu m'a donné n'existent pas pour cette commande")
			return
		}
		c.Run(s, m.Message, given.Param, given.Options)
	}
}
