package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

func handleLotto(s *discordgo.Session, i *discordgo.InteractionCreate) {
	mode := i.ApplicationCommandData().Options[0].StringValue()

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("🎫 lotto 명령어: %s", mode),
		},
	})

	if err != nil {
		log.Println("Failed to respond to ping: ", err)
		replyError(s, i)
		return
	}
}
