package main

import (
	"discord_pterodactyl_connector/config"
	"discord_pterodactyl_connector/handler"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}
	if cfg.DiscordToken == "" || cfg.APIToken == "" || cfg.PterodactylURL == "" || cfg.ServerID == "" {
		log.Fatal("Missing required configuration values. Please check your config.yml file.")
	}

	// Create a new Discord session using the provided token
	dg, err := discordgo.New("Bot " + cfg.DiscordToken)
	if err != nil {
		log.Fatalf("Error creating Discord session: %v", err)
	}
	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		handler.MessageCreate(s, m, cfg)
	})

	err = dg.Open()
	if err != nil {
		log.Fatalf("Error opening Discord session: %v", err)
	}
	fmt.Println("Bot is now running. Press Ctrl+C to exit.")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	dg.Close()
}
