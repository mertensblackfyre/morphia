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
			Sugar.Infof(v.Name)
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

func ReorderRole(s *discordgo.Session, guildID string, role_id string, offset int) error {
	roles, err := s.GuildRoles(guildID)
	if err != nil {
		Sugar.Error(err)
		return err
	}

	var highest_role int = 0
	for _, r := range roles {
		if r.Position > highest_role {
			highest_role = r.Position
		}
	}
	for _, r := range roles {
		if r.ID == role_id {
			r.Position = highest_role - offset
			break
		}
	}

	_, err = s.GuildRoleReorder(guildID, roles)
	if err != nil {
		Sugar.Error(err)
		return err
	}

	return nil
}

func CreateRole(session *discordgo.Session, interaction *discordgo.InteractionCreate) {

	if CheckUserHasRole(interaction.Member.User.ID) == 0 {
		ReplyError(session, interaction, "You already have a role.")
		return
	}

	options := interaction.ApplicationCommandData().Options

	var color int
	var name string = options[0].StringValue()
	var mentionable bool = false

	if len(options) > 1 {
		for _, opt := range options {
			if opt.Name == "color" {
				color = int(opt.IntValue())
			}
		}
	}

	aa := &discordgo.RoleParams{
		Name:        name,
		Color:       &color,
		Mentionable: &mentionable,
	}

	role, err := session.GuildRoleCreate(interaction.GuildID, aa)

	if err != nil {
		ReplyError(session, interaction, err.Error())
		Sugar.Errorln(err)
		return
	}

	err = ReorderRole(session, interaction.GuildID, role.ID,12)

	if err != nil {
		Sugar.Errorln(err)
	}

	err = session.GuildMemberRoleAdd(interaction.GuildID, interaction.Member.User.ID, role.ID)

	if err != nil {
		ReplyError(session, interaction, err.Error())
		Sugar.Errorln(err)
		return
	}

	str := strconv.Itoa(color)
	InsertRoleDB(role.ID, name, interaction.Member.User.ID, str)
	msg := fmt.Sprintf("Role have been created.\n Name: %s\n Color: %s", role.Name, str)
	ReplySuccess(session, interaction, msg)
	Sugar.Infof("Role %s have been created", name)
}

func DeleteRole(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	role := TraverseDB(interaction.Member.User.ID)

	err := session.GuildRoleDelete(interaction.GuildID, role.RoleID)
	if err != nil {
		ReplyError(session, interaction, "Role does not exist")
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
	num, err := strconv.Atoi(r.Color)
	if err != nil {
		Sugar.Error(err)
	}
	var color int = int(num)
	var name string = r.Name

	if len(options) >= 1 {
		for _, opt := range options {
			if opt.Name == "color" {
				color = int(opt.IntValue())
			}
			if opt.Name == "name" {
				name = opt.StringValue()
			}
		}
	} else {
		ReplyError(session, interaction, "No values")
		Sugar.Errorln("No Values")
		return
	}

	role_params := &discordgo.RoleParams{
		Name:  name,
		Color: &color,
	}

	_, err = session.GuildRoleEdit(interaction.GuildID, r.RoleID, role_params)

	if err != nil {
		ReplyError(session, interaction, err.Error())
		Sugar.Errorln(err)
		return
	}

	str := strconv.Itoa(*role_params.Color)
	UpdateRoleDB(name, str, r.RoleID)
	ReplySuccess(session, interaction, "Role Updated")
	Sugar.Infof("Role %s have been updated", name)

}

func GetServerBoosters(session *discordgo.Session) {
	members, err := session.GuildMembers(GUILD_ID, "", 1000)
	if err != nil {
		fmt.Println("Error getting members:", err)
		return
	}
	for _, m := range members {
		if m.PremiumSince != nil {
			fmt.Printf(" Booster: %s (%s)\n", m.User.Username, m.User.ID)
		}
	}
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
