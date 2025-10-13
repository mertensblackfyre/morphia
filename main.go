package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {

	Init()
	Logger()

	Sugar.Infow("App started", "version", "1.0.0")

	dg, err := discordgo.New("Bot " + DISCORD_KEY)

	if err != nil {
		Sugar.Errorw("Error creating Discord dg:", "error", err)
		return
	}

	dg.AddHandler(onInteraction)

	err = dg.Open()
	if err != nil {
		Sugar.Errorw("Error opening connection:", "error", err)
		return
	}

	RegisterCommands(dg)
	Sugar.Infoln("Bot is running. Press CTRL+C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	RemoveCommands(dg)

	defer Sugar.Sync()
	defer DB.Close()
	dg.Close()
}

// Handle slash commands
func onInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()

	switch data.Name {
	case "createrole":
		CreateRole(s, i)
	case "removerole":
		DeleteRole(s, i)
	case "editrole":
		UpdateRole(s, i)

	}
}
