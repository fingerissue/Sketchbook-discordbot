package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	DISCORD_BOT_TOKEN := os.Getenv("DISCORD_BOT_TOKEN")
	if DISCORD_BOT_TOKEN == "" {
		log.Fatal("You need to set DISCORD_BOT_TOKEN")
	}

	Sketchbook, err := discordgo.New("Bot " + DISCORD_BOT_TOKEN)
	if err != nil {
		log.Fatal("Unable to create discord session: ", err)
	}

	err = Sketchbook.Open()
	if err != nil {
		log.Fatal("Unable to connect discord: ", err)
	}
	defer Sketchbook.Close()

	fmt.Println("Bot is now running. Press CTRL-C to exit.")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	fmt.Println("Shutting down...")
}
