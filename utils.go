package main

import (
	"fmt"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func RegisterCommands(session *discordgo.Session) {

	Sugar.Infoln("Registering commands...")
	for _, v := range commands {
		cmd, err := session.ApplicationCommandCreate(session.State.User.ID, GUILD_ID, v)
		if err != nil {
			Sugar.Infof("Cannot create '%v' command: %v\n", v.Name, err)
		} else {
			RegisteredCommands = append(RegisteredCommands, cmd)
		}
	}
}

func RemoveCommands(session *discordgo.Session) {
	for _, cmd := range commands {
		err := session.ApplicationCommandDelete(session.State.User.ID, GUILD_ID, cmd.ID)
		if err != nil {
			Sugar.Errorf("Failed to delete %s: %v\n", cmd.Name, err)
		} else {
			Sugar.Errorf("Deleted /%s\n", cmd.Name)
		}
	}
	session.Close()
}

// Create role and optionally assign it
func CreateRole(session *discordgo.Session, interaction *discordgo.InteractionCreate) {

	if CheckUserHasRole(interaction.Member.User.ID) == 0 {
		ReplyError(session, interaction, "You already have a role")
		return
	}

	options := interaction.ApplicationCommandData().Options
	var color int
	var name string = options[0].StringValue()
	if len(options) > 1 {
		for _, opt := range options {
			if opt.Name == "color" {
				color = int(opt.IntValue())
			}
		}
	}

	aa := &discordgo.RoleParams{
		Name:  name,
		Color: &color,
	}

	role, err := session.GuildRoleCreate(interaction.GuildID, aa)

	if err != nil {
		ReplyError(session, interaction, err.Error())
		Sugar.Errorln(err)
		return
	}

	err = session.GuildMemberRoleAdd(interaction.GuildID, interaction.Member.User.ID, role.ID)

	if err != nil {
		ReplyError(session, interaction, err.Error())
		Sugar.Errorln(err)
		return
	}

	str := strconv.Itoa(color)
	InsertRoleDB(role.ID, name, interaction.Member.User.ID, str)
	ReplySuccess(session, interaction, "Role created")
	Sugar.Infof("Role %s have been created", name)
	return
}

func DeleteRole(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	role := TraverseDB(interaction.Member.User.ID)

	err := session.GuildRoleDelete(interaction.GuildID, role.RoleID)
	if err != nil {
		Sugar.Errorln(err)
		return
	}
	RemoveRoleDB(role.RoleID, role.UserID)
	Sugar.Infof("Role have been removed")
	ReplySuccess(session, interaction, "Your role have been deleted")
}

func UpdateRole(session *discordgo.Session, interaction *discordgo.InteractionCreate) {

	r := TraverseDB(interaction.Member.User.ID)
	options := interaction.ApplicationCommandData().Options

	var color int
	var name string

	if len(options) >= 1 {
		for _, opt := range options {
			if opt.Name == "color" {
				color = int(opt.IntValue())
			}
			if  opt.Name == "name"{
				name = opt.StringValue()
			}
		}
	} else {
		ReplyError(session, interaction, "No values")
		Sugar.Errorln("No Values")
		return
	}

	fmt.Println(name)
	fmt.Println(color)
	role_params := &discordgo.RoleParams{
		Name:  name,
		Color: &color,
	}

	_, err := session.GuildRoleEdit(interaction.GuildID, r.RoleID, role_params)

	if err != nil {
		ReplyError(session, interaction, err.Error())
		Sugar.Errorln(err)
		return
	}

	str := strconv.Itoa(color)
	UpdateRoleDB(name, str, r.RoleID)
	ReplySuccess(session, interaction, "Role Updated")
	Sugar.Infof("Role %s have been updated", name)
	return

}

func ReplyError(s *discordgo.Session, i *discordgo.InteractionCreate, msg string) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: msg,
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}

func ReplySuccess(s *discordgo.Session, i *discordgo.InteractionCreate, msg string) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: msg,
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}
