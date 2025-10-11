package main

import (
	"github.com/bwmarrin/discordgo"
)

// Define the slash commands
var commands = []*discordgo.ApplicationCommand{

	{
		Name:        "createrole",
		Description: "Creates a new role in this server",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "name",
				Description: "Name of the new role",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "color",
				Description: "Color of the role (integer hex, e.g. 16711680 for red)",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionBoolean,
				Name:        "assignme",
				Description: "Whether to assign the role to you automatically",
				Required:    false,
			},
		},
	},
}
