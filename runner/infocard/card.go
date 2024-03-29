package infocard

import (
	"annyobot/command/model"
	"encoding/json"
	"io"
	"net/http"

	"strings"

	"github.com/bwmarrin/discordgo"
)

var basicURI = "https://db.ygoprodeck.com/api/v7/cardinfo.php"
var fr = "language=fr"

type JSON struct {
	Card []Card `json:"data"`
}

type Card struct {
	Name   string    `json:"name"`
	Images []CardImg `json:"card_images"`
}

type CardImg struct {
	Id       int    `json:"id"`
	Url      string `json:"image_url"`
	UrlSmall string `json:"image_url_small"`
}

func PrintCard(session *discordgo.Session, message *discordgo.Message, param string, options []model.Option) {
	paramSplit := strings.Split(param, " ")
	finalURI := basicURI + "?name="
	for i, s := range paramSplit {
		if i != len(paramSplit)-1 {
			finalURI += s + "%20"
		} else {
			finalURI += s
		}
	}
	if options != nil {
		finalURI += "&" + fr
	}

	response, err := http.Get(finalURI)
	if err != nil {
		_, _ = session.ChannelMessageSend(message.ChannelID, "J'ai eu un problème en allant chercher ta carte")
	}
	resBody, _ := io.ReadAll(response.Body)

	var c JSON
	_ = json.Unmarshal(resBody, &c)

	if len(c.Card) == 0 {
		_, _ = session.ChannelMessageSend(message.ChannelID, "Je n'est pas trouvé ta carte")
	} else {
		_, _ = session.ChannelMessageSendEmbed(message.ChannelID, &discordgo.MessageEmbed{
			Description: c.Card[0].Name,
			Color:       0x000000,
			Image: &discordgo.MessageEmbedImage{
				URL: c.Card[0].Images[0].Url,
			},
		})
	}
}
