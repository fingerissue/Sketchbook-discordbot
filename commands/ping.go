package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func handlePing(s *discordgo.Session, i *discordgo.InteractionCreate) {
	ping := s.HeartbeatLatency().Milliseconds()
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("🏓 Pong! %dms", ping),
		},
	})
}
