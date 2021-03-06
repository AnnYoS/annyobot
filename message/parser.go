package message

import (
	"annyobot/command/model"

	"fmt"
	"strings"
)

func TransformToCommand(msg string) (*model.Command, error) {
	lastChar := "!"
	for _, s := range msg {
		if strings.HasPrefix(string(s), "-") && lastChar != " " {
			return nil, fmt.Errorf("Bad option")
		}
		lastChar = string(s)
	}

	var c model.Command
	split := strings.Split(msg, " ")
	c.Command = split[0]
	c.Param = ""

	find := false
	for _, s := range split {
		if strings.HasPrefix(s, "'") && strings.HasSuffix(s, "'") {
			c.Param += s[1 : len(s)-1]
			find = true
			break
		} else if strings.HasPrefix(s, "'") && !strings.HasSuffix(s, "'") {
			c.Param += s[1:]
			find = true
		} else if find && !strings.HasSuffix(s, "'") {
			c.Param += " " + s
		} else if find && strings.HasSuffix(s, "'") {
			c.Param += " " + s[:len(s)-1]
			break
		}
	}

	for _, s := range split {
		if strings.HasPrefix(s, "-") {
			c.Options = append(c.Options, model.Option{Opt: s, Desc: ""})
		}
	}
	return &c, nil
}

func VerifyOptions(opts []model.Option, theoric []model.Option) bool {
	var res = false
	for _, s := range opts {
		res = false
		for _, t := range theoric {
			if s.Opt == t.Opt {
				res = true
			}
		}
		if !res {
			return false
		}
	}
	return true
}
