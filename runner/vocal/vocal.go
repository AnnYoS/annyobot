package vocal

import (
	"annyobot/command/model"
	"github.com/bwmarrin/discordgo"
)

func JoinVocal(session *discordgo.Session, message *discordgo.Message, param string, options []model.Option) {
	g, _ := session.State.Guild(message.GuildID)
	u, _ := session.GuildMember(message.GuildID, message.Author.ID)

	for _, state := range g.VoiceStates {
		if state.UserID == u.User.ID {
			channel, _ := session.State.Channel(state.ChannelID)
			_, _ = session.ChannelVoiceJoin(message.GuildID, channel.ID, true, false)
		}
	}
}

func ExitVocal(session *discordgo.Session, message *discordgo.Message, param string, options []model.Option) {

}
