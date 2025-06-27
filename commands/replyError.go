package commands

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func replyError(s *discordgo.Session, i *discordgo.InteractionCreate) {
	_, err := s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
		Content: "⚠️ 응답 중 오류가 발생했습니다.",
	})

	if err != nil {
		log.Println("FollowupMessageCreate error: ", err)
	}
}
