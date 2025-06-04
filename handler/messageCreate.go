package handler

import (
	"discord_pterodactyl_connector/config"
	"discord_pterodactyl_connector/pterodactyl"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate, config *config.Config) {
	if m.Author.ID == s.State.User.ID {
		return // Ignore messages from the bot itself
	}

	content := m.Content
	prefix := config.CommandPrefix

	switch {
	case startsWith(content, prefix+"start"):
		handleStart(s, m, config)
	case startsWith(content, prefix+"stop"):
		handleStop(s, m, config)
	case startsWith(content, prefix+"status"):
		handleStatus(s, m, config)
	case startsWith(content, prefix+"restart"):
		handleRestart(s, m, config)
	case startsWith(content, prefix+"kill"):
		handleKill(s, m, config)
	case startsWith(content, prefix+"help"):
		handleHelp(s, m, config)
	}
}

func handleStart(s *discordgo.Session, m *discordgo.MessageCreate, config *config.Config) {
	status := pterodactyl.GetStatus(config)
	log.Printf("Server status: %s", status)
	if status == "running" {
		s.ChannelMessageSend(m.ChannelID, "Server is already running.")
		return
	}
	if status == "starting" {
		s.ChannelMessageSend(m.ChannelID, "Server is already starting.")
		return
	}
	if status == "stopping" {
		s.ChannelMessageSend(m.ChannelID, "Server is currently stopping. Please wait until it is stopped before starting it again.")
		return
	}
	if status == "offline" {
		s.ChannelMessageSend(m.ChannelID, "Server is offline. Starting it now...")
	} else if status == "unknown" {
		s.ChannelMessageSend(m.ChannelID, "Server status is unknown. Attempting to start it...")
	}
	pterodactyl.SendPowerSignal("start", config)
	s.ChannelMessageSend(m.ChannelID, "Server start command sent. Please wait for the server to start.")
}

func handleStop(s *discordgo.Session, m *discordgo.MessageCreate, config *config.Config) {
	status := pterodactyl.GetStatus(config)
	if status == "offline" {
		s.ChannelMessageSend(m.ChannelID, "Server is already offline.")
		return
	}
	if status == "stopping" {
		s.ChannelMessageSend(m.ChannelID, "Server is already stopping.")
		return
	}
	if status == "starting" {
		s.ChannelMessageSend(m.ChannelID, "Server is currently starting. Please wait until it is started before stopping it.")
		return
	}
	if status == "running" {
		s.ChannelMessageSend(m.ChannelID, "Server is running. Stopping it now...")
		pterodactyl.SendPowerSignal("stop", config)
		s.ChannelMessageSend(m.ChannelID, "Server stop command sent. Please wait for the server to stop.")
		return
	} else if status == "unknown" {
		s.ChannelMessageSend(m.ChannelID, "Server status is unknown. Attempting to stop it...")
		pterodactyl.SendPowerSignal("stop", config)
		return
	}
}

func handleStatus(s *discordgo.Session, m *discordgo.MessageCreate, config *config.Config) {
	status := pterodactyl.GetStatus(config)
	message := fmt.Sprintf("Server status: %s", status)
	s.ChannelMessageSend(m.ChannelID, message)
}

func handleRestart(s *discordgo.Session, m *discordgo.MessageCreate, config *config.Config) {
	status := pterodactyl.GetStatus(config)
	if status == "running" {
		pterodactyl.SendPowerSignal("restart", config)
		s.ChannelMessageSend(m.ChannelID, "Server restart command sent. Please wait for the server to restart.")
		return
	} else if status == "stopping" {
		s.ChannelMessageSend(m.ChannelID, "Server is currently stopping. Please wait until it is stopped before restarting it.")
		return
	} else if status == "starting" {
		s.ChannelMessageSend(m.ChannelID, "Server is currently starting. Please wait until it is started before restarting it.")
		return
	} else if status == "offline" {
		s.ChannelMessageSend(m.ChannelID, "Server is offline. Starting it now...")
		pterodactyl.SendPowerSignal("start", config)
		return
	} else if status == "unknown" {
		s.ChannelMessageSend(m.ChannelID, "Server status is unknown. Attempting to restart it...")
		pterodactyl.SendPowerSignal("restart", config)
		return
	}
}

func handleKill(s *discordgo.Session, m *discordgo.MessageCreate, config *config.Config) {
	err := pterodactyl.SendPowerSignal("kill", config)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Error sending kill command: "+err.Error())
	}
}

func handleHelp(s *discordgo.Session, m *discordgo.MessageCreate, config *config.Config) {
	helpMessage := fmt.Sprintf("Available commands:\n"+
		"%sstart - Start the server\n"+
		"%sstop - Stop the server\n"+
		"%sstatus - Get the server status\n"+
		"%srestart - Restart the server\n"+
		"%skill - Kill the server\n"+
		"%shelp - Show this help message",
		config.CommandPrefix, config.CommandPrefix, config.CommandPrefix, config.CommandPrefix, config.CommandPrefix, config.CommandPrefix)
	s.ChannelMessageSend(m.ChannelID, helpMessage)
}

func startsWith(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}
