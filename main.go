package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	tokenDiscord := os.Getenv("TOKEN")

	discord, err := discordgo.New("Bot " + tokenDiscord)
	if err != nil {
		log.Fatal("Error creating discord sesion")
	}

	err = discord.Open()
	if err != nil {
		fmt.Println("Error opening connection", err)
	}

	fmt.Println("Bot Runing!")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.Close()
}
