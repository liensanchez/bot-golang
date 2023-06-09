package commands

import (
	"bot-golang/services"

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

	if m.Content == "!frase" {
		// Send a "Pong!" response
		frases := services.GetFrase()

		for _, f := range frases {
			s.ChannelMessageSend(m.ChannelID, f.Text)
		}
	}

	if m.Content == "!chiste" {
		// Send a "Pong!" response
		frases := services.GetChiste()

		for _, f := range frases {
			s.ChannelMessageSend(m.ChannelID, f.Text)
		}
	}

	if m.Content == "!refran" {
		// Send a "Pong!" response
		frases := services.GetRefran()

		for _, f := range frases {
			s.ChannelMessageSend(m.ChannelID, f.Text)
		}
	}
}
