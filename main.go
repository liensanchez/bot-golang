package main

import (
	"bot-golang/commands"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {

	_ = godotenv.Load()
	TOKEN := os.Getenv("TOKEN")

	dg, err := discordgo.New("Bot " + TOKEN)
	if err != nil {
		fmt.Println(err)
		return
	}

	dg.AddHandler(ready)
	dg.AddHandler(commands.MessageCreate)

	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection", err)
		return
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	fmt.Println("Bot iniciado")
}
