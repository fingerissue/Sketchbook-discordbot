package commands

import (
	"database/sql"
	"github.com/bwmarrin/discordgo"
)

var DB *sql.DB

func OnInteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.ApplicationCommandData().Name {
	case "ping":
		handlePing(s, i)
	case "lotto":
		handleLotto(s, i)
	}
}
