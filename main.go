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

	//empezamos cargando el .env con nuestras variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	tokenDiscord := os.Getenv("TOKEN")

	//Seteamos nuestro Token para el bot
	discord, err := discordgo.New("Bot " + tokenDiscord)
	if err != nil {
		log.Fatal("Error creating discord sesion")
	}

	err = discord.Open()
	if err != nil {
		fmt.Println("Error opening connection", err)
		return
	}

	fmt.Println("Bot Runing!")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	discord.Close()
}
