package commands

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func replyError(s *discordgo.Session, i *discordgo.InteractionCreate, msg string) {
	_, err := s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
		Content: msg,
	})

	if err != nil {
		log.Println("FollowupMessageCreate error: ", err)
	}
}
