package commands

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if strings.HasPrefix(m.Content, ":milei") {
		fmt.Println("Hello Milei")
		s.ChannelMessageSend(m.ChannelID, "hellomilei")
	}
}
