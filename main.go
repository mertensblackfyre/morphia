package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			Sugar.Error(r)
		}
	}()

	Init()
	Logger()

	Sugar.Infow("App started", "version", "1.0.0")

	dg, err := discordgo.New("Bot " + DISCORD_KEY)

	if err != nil {
		Sugar.Errorw("Error creating Discord dg:", "error", err)
		return
	}

	go dg.AddHandler(onInteraction)

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

	GetPremuimUsers(s)
	switch data.Name {
	case "createrole":
		go CreateRole(s, i)
	case "removerole":
		go DeleteRole(s, i)
	case "editrole":
		go UpdateRole(s, i)

	}
}
