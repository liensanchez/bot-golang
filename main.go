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

// Define the command structure
type PingCommand struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Register the command with the Discord API
func registerCommands(s *discordgo.Session) error {
	guildID := os.Getenv("GUILD_ID")

	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "ping",
			Description: "Responds with Pong!",
		},
	}

	_, err := s.ApplicationCommandBulkOverwrite(guildID, "", commands)
	if err != nil {
		return err
	}

	return nil
}

// Handle the slash command
func handleCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.ApplicationCommandData().Name {
	case "ping":
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Pong!",
			},
		})
	}
}

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

	// Register the slash commands
	err = registerCommands(discord)
	if err != nil {
		log.Fatal("Error registering commands: ", err)
	}

	// Add the command handler
	discord.AddHandler(handleCommand)

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
