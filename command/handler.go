package command

import (
	"annyobot/runner/basic"
	"annyobot/runner/infocard"
)

var Association map[string]Command

func FindCommand(msg string) *Command {
	c, found := Association[msg]
	if !found {
		return nil
	}
	return &c
}

func Associate() {
	Association = make(map[string]Command)
	Association["!hello"] = Command{Command: "!hello", Desc: "Say hello, nothing more", Param: "",
		Options: []Option{{Opt: "-a", Desc: "Say hello to everyone"}, {Opt: "-m", Desc: "Say hello to yourself"}}, Run: basic.Hello}

	Association["!help"] = Command{Command: "!help", Desc: "Giving help about my commands", Param: "None or command you want to know about (example for !hello, give hello)",
		Options: []Option{}, Run: basic.Help}

	Association["!start"] = Command{Command: "!start", Desc: "Give the basic about the bot", Param: "",
		Options: []Option{}, Run: basic.GetStart}

	Association["!infocard"] = Command{Command: "!infocard", Desc: "I print Yu-Gi-Oh infocard you want !", Param: "Name of the infocard",
		Options: []Option{{Opt: "-f", Desc: "Name in french"}}, Run: infocard.PrintCard}
}
