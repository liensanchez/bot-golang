package commands

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Println("Hello Milei")
	if strings.HasPrefix(m.Content, ":milei") {
		fmt.Println("Hello Milei2")
		s.ChannelMessageSend(m.ChannelID, "hellomilei")
	}
}
