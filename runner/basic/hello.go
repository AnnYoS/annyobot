package basic

import (
	"annyobot/command/model"

	"github.com/bwmarrin/discordgo"
)

func GetStart(session *discordgo.Session, message *discordgo.Message, param string, options []model.Option) {
	msg := "```Salut à toi mon nouvel utilisateur, je suis AnnYo V.Bot mais tu peux m'appeler AnnYo !\n" +
		"Je tire mon nom de mon créateur AnnYo ou BitConneeeeect#8080 si tu préfères. Je suis fait en Go si ça t'intéresse de savoir !\n" +
		"Bon assez parlé de moi, tu es la pour savoir comment te servir de moi. Je te conseille la commande !help, elle" +
		"t'offrira toutes les commandes que je possède. Et dernière petite info, si tu souhaites voir mon code, c'est par ici:\n\n" +
		"https://github.com/AnnYoS/annyobot```"
	_, _ = session.ChannelMessageSend(message.ChannelID, msg)
}

func Hello(session *discordgo.Session, message *discordgo.Message, param string, options []model.Option) {
	if len(options) > 1 {
		_, _ = session.ChannelMessageSend(message.ChannelID, "Oula tu m'en demmande trop")

	} else {
		if options == nil {
			_, _ = session.ChannelMessageSend(message.ChannelID, "Salut !")
		} else if options[0].Opt == "-a" {
			_, _ = session.ChannelMessageSend(message.ChannelID, "Salut @everyone")
		} else if options[0].Opt == "-m" {
			_, _ = session.ChannelMessageSend(message.ChannelID, "Salut "+message.Author.Mention())
		}
	}
}
