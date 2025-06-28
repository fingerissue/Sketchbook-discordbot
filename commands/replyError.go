package commands

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func replyErrorInteraction(s *discordgo.Session, i *discordgo.InteractionCreate, msg string) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: msg,
		},
	})

	if err != nil {
		log.Println("While error message sending error: ", err)
	}
}

func replyErrorFollowup(s *discordgo.Session, i *discordgo.InteractionCreate, msg string) {
	_, err := s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
		Content: msg,
	})

	if err != nil {
		log.Println("While error message sending error: ", err)
	}
}
