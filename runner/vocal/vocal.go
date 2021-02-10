package vocal

import (
	"annyobot/command/model"
	"github.com/bwmarrin/discordgo"
)

var voiceConnection *discordgo.VoiceConnection

func JoinVocal(session *discordgo.Session, message *discordgo.Message, param string, options []model.Option) {
	g, _ := session.State.Guild(message.GuildID)
	u, _ := session.GuildMember(message.GuildID, message.Author.ID)

	var connection *discordgo.VoiceConnection
	for _, state := range g.VoiceStates {
		if state.UserID == u.User.ID {
			channel, _ := session.State.Channel(state.ChannelID)
			connection, _ = session.ChannelVoiceJoin(message.GuildID, channel.ID, true, false)
		}
	}
	voiceConnection = connection
}

func ExitVocal(session *discordgo.Session, message *discordgo.Message, param string, options []model.Option) {
	if voiceConnection != nil {
		voiceConnection.Close()
		err := voiceConnection.Disconnect()
		if err != nil {
			_, _ = session.ChannelMessageSend(message.ChannelID, "Je ne peux pas me déconnecter si je ne suis même pas en vocal")
		}
	}
}
