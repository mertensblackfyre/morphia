package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func RegisterCommands(session *discordgo.Session) {

	fmt.Println("Registering commands...")
	for _, v := range commands {
		cmd, err := session.ApplicationCommandCreate(session.State.User.ID, GUILD_ID, v)
		if err != nil {
			fmt.Printf("Cannot create '%v' command: %v\n", v.Name, err)
		} else {
			RegisteredCommands = append(RegisteredCommands, cmd)
		}
	}
}

func RemoveCommands(session *discordgo.Session) {
	for _, cmd := range commands {
		err := session.ApplicationCommandDelete(session.State.User.ID, GUILD_ID, cmd.ID)
		if err != nil {
			fmt.Printf("Failed to delete %s: %v\n", cmd.Name, err)
		} else {
			fmt.Printf("Deleted /%s\n", cmd.Name)
		}
	}
	session.Close()
	fmt.Println("👋 Bot shut down.")
}

// Create role and optionally assign it
func CreateRole(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	options := interaction.ApplicationCommandData().Options
	roleName := options[0].StringValue()

	var color int

	if len(options) > 1 {
		for _, opt := range options {
			if opt.Name == "color" {
				color = int(opt.IntValue())
			}

		}
	}

	role, err := session.GuildRoleCreate(interaction.GuildID)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	/*if assignMe {
	err := session.GuildMemberRoleAdd(interaction.GuildID, interaction.Member.User.ID, role.ID)
	if err != nil {
		fmt.Println(err)
		return
	}
	*/
	return

}
