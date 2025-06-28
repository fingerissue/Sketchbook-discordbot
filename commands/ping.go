package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

func handlePing(s *discordgo.Session, i *discordgo.InteractionCreate) {
	ping := s.HeartbeatLatency().Milliseconds()
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("🏓 Pong! %dms", ping),
		},
	})

	if err != nil {
		log.Println("Failed to respond to ping: ", err)
		replyError(s, i, "⚠️ 응답 중 오류가 발생했습니다.")
		return
	}
}
