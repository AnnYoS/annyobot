package command

import (
	"annyobot/command/model"
	"annyobot/runner/basic"
	"annyobot/runner/infocard"
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
	model.Association["!hello"] = model.Command{Command: "!hello", Desc: "Say hello, nothing more", Param: "",
		Options: []model.Option{{Opt: "-a", Desc: "Say hello to everyone"}, {Opt: "-m", Desc: "Say hello to yourself"}}, Run: basic.Hello}

	model.Association["!help"] = model.Command{Command: "!help", Desc: "Giving help about my commands", Param: "None or command you want to know about (example for !hello, give hello)",
		Options: []model.Option{}, Run: Help}

	model.Association["!start"] = model.Command{Command: "!start", Desc: "Give the basic about the bot", Param: "",
		Options: []model.Option{}, Run: basic.GetStart}

	model.Association["!infocard"] = model.Command{Command: "!infocard", Desc: "I print Yu-Gi-Oh infocard you want !", Param: "Name of the infocard",
		Options: []model.Option{{Opt: "-f", Desc: "Name in french"}}, Run: infocard.PrintCard}
}
