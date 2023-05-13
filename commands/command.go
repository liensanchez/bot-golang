package commands

import (
	"os"

	"github.com/bwmarrin/discordgo"
)

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
