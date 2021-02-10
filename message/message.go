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
			_, _ = s.ChannelMessageSend(m.ChannelID, "Careful options !")
			return
		}
		c := command.FindCommand(given.Command)
		if c == nil {
			_, _ = s.ChannelMessageSend(m.ChannelID, "Invalid command, try !help to know how to use me :smile:")
			return
		}

		correct := VerifyOptions(given.Options, c.Options)
		if !correct {
			_, _ = s.ChannelMessageSend(m.ChannelID, "Unknown option !")
			return
		}
		c.Run(s, m.Message, given.Param, given.Options)
	}
}