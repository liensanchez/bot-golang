package commands

import (
	"github.com/bwmarrin/discordgo"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	// Check if the message content is "!ping"
	if m.Content == "!ping" {
		// Send a "Pong!" response
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

}
