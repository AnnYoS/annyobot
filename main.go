package main

import (
	"annyobot/command"
	"annyobot/message"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var Token string

func main() {
	initConsole()
	initDiscordSession()
}

func initConsole() {
	fmt.Println(`     
	 _             __   __     __     ______        _   
    / \   _ __  _ _\ \ / /__   \ \   / / __ )  ___ | |_ 
   / _ \ | '_ \| '_ \ V / _ \   \ \ / /|  _ \ / _ \| __|
  / ___ \| | | | | | | | (_) |   \ V / | |_) | (_) | |_ 
 /_/   \_\_| |_|_| |_|_|\___/     \_(_)|____/ \___/ \__|`)
	fmt.Println("")
}

func initDiscordSession() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err.Error())
	}
	if Token = os.Getenv("TOKEN"); Token == "" {
		log.Fatalln("Bad token")
	}

	session, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Fatalln("Cannot create Discord session")
	}
	err = session.Open()
	if err != nil {
		log.Fatalln("Cannot connect to Discord")
	}

	command.Associate()
	session.AddHandler(message.ReceiveMessage)

	fmt.Println("Discord bot is now running, waiting to exit")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	_ = session.Close()
}