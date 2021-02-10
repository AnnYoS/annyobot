package model

import "github.com/bwmarrin/discordgo"

var Association map[string]Command

type Command struct {
	Command string
	Desc    string
	Param   string
	Options []Option
	Run     ActionFunction
}

type Option struct {
	Opt  string
	Desc string
}

type ActionFunction func(*discordgo.Session, *discordgo.Message, string, []Option)
