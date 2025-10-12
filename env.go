package main

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	TURSO_DATABASE_URL string
	TURSO_AUTH_TOKEN string
	DISCORD_KEY string
	GUILD_ID    = "1266891220043563018"
	RegisteredCommands []*discordgo.ApplicationCommand
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	
	TURSO_DATABASE_URL = os.Getenv("TURSO_DATABASE_URL")
	TURSO_AUTH_TOKEN = os.Getenv("TURSO_AUTH_TOKEN")
	DISCORD_KEY = os.Getenv("DISCORD_KEY")
}
