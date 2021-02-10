package command

import (
	"annyobot/command/model"
	"annyobot/runner/basic"
	"annyobot/runner/infocard"
	"annyobot/runner/vocal"
)

func FindCommand(msg string) *model.Command {
	c, found := model.Association[msg]
	if !found {
		return nil
	}
	return &c
}

func Associate() {
	model.Association = make(map[string]model.Command)
	model.Association["!hello"] = model.Command{Command: "!hello", Desc: "Je salut comme il se doit", Param: "",
		Options: []model.Option{{Opt: "-a", Desc: "Je salut tout le monde"}, {Opt: "-m", Desc: "Je te dit bonjour qu'à toi"}}, Run: basic.Hello}

	model.Association["!help"] = model.Command{Command: "!help", Desc: "Je t'aide sur mon utilisation", Param: "Aucun ou bien la commande qui t'interesse (exemple: pour !hello, donne moi hello)",
		Options: []model.Option{}, Run: Help}

	model.Association["!start"] = model.Command{Command: "!start", Desc: "Je m'introduis", Param: "",
		Options: []model.Option{}, Run: basic.GetStart}

	model.Association["!infocard"] = model.Command{Command: "!infocard", Desc: "Je t'affiche une carte Yu Gi Oh car j'aime ce jeu !", Param: "Nom de la carte (par defaut en anglais)",
		Options: []model.Option{{Opt: "-f", Desc: "Si tu m'a donné le nom en français"}}, Run: infocard.PrintCard}

	model.Association["!join"] = model.Command{Command: "!join", Desc: "Je rejoint le vocal où tu es", Param: "Aucun",
		Options: []model.Option{}, Run: vocal.JoinVocal}

	model.Association["!voicekick"] = model.Command{Command: "!voicekick", Desc: "Si tu souhaites me kick du vocal", Param: "Aucun",
		Options: []model.Option{}, Run: vocal.ExitVocal}
}
