package main

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	DISCORD_KEY string
	GUILD_ID    = "1266891220043563018"
	RegisteredCommands []*discordgo.ApplicationCommand
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	DISCORD_KEY = os.Getenv("DISCORD_KEY")
}
