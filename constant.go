package main

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

type Role struct {
	ID        int       `db:"id"`
	UserID    string    `db:"user_id"`
	RoleID    string    `db:"role_id"`
	Name      string    `db:"name"`
	Color     string    `db:"color"`
	CreatedAt time.Time `db:"created_at"`
}

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
		},
	},
	{
		Name:        "editrole",
		Description: "Edit your role in this server",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "name",
				Description: "New name of the role",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "color",
				Description: "The new color of the role (integer hex, e.g. 0x16711680 for red)",
				Required:    false,
			},
		},
	},

	{
		Name:        "removerole",
		Description: "Remove your role",
	},
}
