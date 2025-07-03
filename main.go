package main

import (
	"Sketchbook/commands"
	"database/sql"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
)

import _ "github.com/go-sql-driver/mysql"

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	DiscordBotToken := os.Getenv("DISCORD_BOT_TOKEN")
	if DiscordBotToken == "" {
		log.Fatal("You need to set DISCORD_BOT_TOKEN")
	}

	DatabaseAddress := os.Getenv("DB_ADDR")
	if DatabaseAddress == "" {
		log.Fatal("You need to set DB_ADDR")
	}

	DatabaseID := os.Getenv("DB_ID")
	if DatabaseID == "" {
		log.Fatal("You need to set DB_ID")
	}

	DatabasePW := os.Getenv("DB_PW")
	if DatabasePW == "" {
		log.Fatal("You need to set DB_PW")
	}

	DatabaseNAME := os.Getenv("DB_NAME")
	if DatabaseNAME == "" {
		log.Fatal("You need to set DB_NAME")
	}

	Sketchbook, err := discordgo.New("Bot " + DiscordBotToken)
	if err != nil {
		log.Fatal("Unable to create discord session: ", err)
	}

	Sketchbook.AddHandler(commands.OnInteractionCreate)

	err = Sketchbook.Open()
	if err != nil {
		log.Fatal("Unable to connect discord: ", err)
	}
	defer func() {
		if err := Sketchbook.Close(); err != nil {
			log.Fatal("Unable to close discord session: ", err)
		}
	}()

	command := []*discordgo.ApplicationCommand{
		{
			Name:        "ping",
			Description: "Replies with pong.",
		},
		{
			Name:        "lotto",
			Description: "Challenge your luck.",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "mode",
					Description: "수동/자동을 선택하세요.",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    true,
					Choices: []*discordgo.ApplicationCommandOptionChoice{
						{Name: "수동", Value: "수동"},
						{Name: "자동", Value: "자동"},
					},
				},
				{
					Name:        "numbers",
					Description: "1부터 45까지 중 6개의 정수를 중복없이 공백으로 구분하여 입력하세요.",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    false,
				},
			},
		},
	}

	for _, cmd := range command {
		_, err = Sketchbook.ApplicationCommandCreate(Sketchbook.State.User.ID, "", cmd)
		if err != nil {
			log.Fatal("Unable to create application command: ", err)
		}
	}

	DB, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", DatabaseID, DatabasePW, DatabaseAddress, DatabaseNAME))
	if err != nil {
		log.Fatal("Unable to connect database: ", err)
	}
	defer func() {
		if err := DB.Close(); err != nil {
			log.Fatal("Unable to close database: ", err)
		}
	}()

	fmt.Println("Bot is now running. Press CTRL-C to exit.")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	fmt.Println("Shutting down...")
}
