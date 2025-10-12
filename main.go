package main

import (
	//"fmt"
	//"os"
	//"os/signal"
	//"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {

	run()
	/*
	dg, err := discordgo.New("Bot " + DISCORD_KEY)
	if err != nil {
		fmt.Println("Error creating Discord dg:", err)
		return
	}

	dg.AddHandler(onInteraction)


	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection:", err)
		return
	}

	RegisterCommands(dg)
	fmt.Println("Bot is running. Press CTRL+C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	RemoveCommands(dg)

	dg.Close()
	*/
}

// Handle slash commands
func onInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()

	switch data.Name {
	case "createrole":
		CreateRole(s, i)
	}
}
