package handler

import (
	"discord_pterodactyl_connector/config"
	"discord_pterodactyl_connector/pterodactyl"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate, config *config.Config) {

	if m.Author.ID == s.State.User.ID {
		return // Ignore messages from the bot itself
	}

	if startsWith(m.Content, config.CommandPrefix+"start") {
		pterodactyl.StartServer(config)
		_, _ = s.ChannelMessageSend(m.ChannelID, "Server start command sent.")
		return
	}

	if startsWith(m.Content, config.CommandPrefix+"stop") {
		pterodactyl.StopServer(config)
		_, _ = s.ChannelMessageSend(m.ChannelID, "Server stop command sent.")
		return
	}
	if startsWith(m.Content, config.CommandPrefix+"status") {
		status := pterodactyl.GetStatus(config)
		_, _ = s.ChannelMessageSend(m.ChannelID, status)
		return
	}
	if startsWith(m.Content, config.CommandPrefix+"restart") {
		pterodactyl.RestartServer(config)
		_, _ = s.ChannelMessageSend(m.ChannelID, "Server restart command sent.")
		return
	}
	if startsWith(m.Content, config.CommandPrefix+"kill") {
		pterodactyl.KillServer(config)
		_, _ = s.ChannelMessageSend(m.ChannelID, "Server kill command sent.")
		return
	}

	if startsWith(m.Content, config.CommandPrefix+"help") {
		helpMessage := fmt.Sprintf("Available commands:\n"+
			"%sstart - Start the server\n"+
			"%sstop - Stop the server\n"+
			"%sstatus - Get the server status\n"+
			"%srestart - Restart the server\n"+
			"%skill - Kill the server\n"+
			"%shelp - Show this help message",
			config.CommandPrefix, config.CommandPrefix, config.CommandPrefix, config.CommandPrefix, config.CommandPrefix, config.CommandPrefix)
		_, _ = s.ChannelMessageSend(m.ChannelID, helpMessage)
		return
	}
}

func startsWith(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}
