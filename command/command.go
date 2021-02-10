package command

import "github.com/bwmarrin/discordgo"

type Command struct {
	Command   string
	Desc 	  string
	Param 	  string
	Options   []Option
	Run 	  ActionFunction
}

type Option struct {
	Opt    string
	Desc   string
}

type ActionFunction func(*discordgo.Session, *discordgo.Message, string, []Option)
