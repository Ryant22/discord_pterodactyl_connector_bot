package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"gopkg.in/yaml.v2"
)

type Config struct {
	DiscordToken   string `yaml:"discord_token"`
	CommandPrefix  string `yaml:"command_prefix"`
	APIToken       string `yaml:"api_token"`
	PterodactylURL string `yaml:"pterodactyl_url"`
	ServerID       string `yaml:"server_id"`
}

var config Config

func loadConfig() {
	configFile, err := os.ReadFile("config.yml")
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
		return
	}
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatalf("Error parsing config file: %v", err)
		return
	}

}

func main() {
	loadConfig()
	if config.DiscordToken == "" || config.APIToken == "" || config.PterodactylURL == "" || config.ServerID == "" {
		log.Fatal("Missing required configuration values. Please check your config.yml file.")
	}
	// Create a new Discord session using the provided token
	dg, err := discordgo.New("Bot " + config.DiscordToken)
	if err != nil {
		log.Fatalf("Error creating Discord session: %v", err)
	}

	dg.AddHandler(messageCreate)

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

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return // Ignore messages from the bot itself
	}

	if startsWith(m.Content, config.CommandPrefix+"start") {
		startServer(&config)
		_, _ = s.ChannelMessageSend(m.ChannelID, "Server start command sent.")
		return
	}

	if startsWith(m.Content, config.CommandPrefix+"stop") {
		stopServer(&config)
		_, _ = s.ChannelMessageSend(m.ChannelID, "Server stop command sent.")
		return
	}
	if startsWith(m.Content, config.CommandPrefix+"status") {
		status := getStatus(&config)
		_, _ = s.ChannelMessageSend(m.ChannelID, status)
		return
	}
	if startsWith(m.Content, config.CommandPrefix+"restart") {
		restartServer(&config)
		_, _ = s.ChannelMessageSend(m.ChannelID, "Server restart command sent.")
		return
	}
	if startsWith(m.Content, config.CommandPrefix+"kill") {
		killServer(&config)
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

func startServer(config *Config) {
	url := fmt.Sprintf("%sservers/%s/power", config.PterodactylURL, config.ServerID)
	payload := map[string]string{"signal": "start"}
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %v", err)
		return
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+config.APIToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		log.Printf("Error response from server: %s", body)
		return
	}
	log.Println("Server started successfully")
}

func stopServer(config *Config) {
	url := fmt.Sprintf("%sservers/%s/power", config.PterodactylURL, config.ServerID)
	payload := map[string]string{"signal": "stop"}
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %v", err)
		return
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+config.APIToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		log.Printf("Error response from server: %s", body)
		return
	}
	log.Println("Server stopped successfully")
}

func restartServer(config *Config) {
	url := fmt.Sprintf("%sservers/%s/power", config.PterodactylURL, config.ServerID)
	payload := map[string]string{"signal": "restart"}
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %v", err)
		return
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+config.APIToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		log.Printf("Error response from server: %s", body)
		return
	}
	log.Println("Server restarted successfully")
}

func killServer(config *Config) {
	url := fmt.Sprintf("%sservers/%s/power", config.PterodactylURL, config.ServerID)
	payload := map[string]string{"signal": "kill"}
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %v", err)
		return
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+config.APIToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		log.Printf("Error response from server: %s", body)
		return
	}
	log.Println("Server killed successfully")
}

func getStatus(config *Config) string {
	url := fmt.Sprintf("%sservers/%s/resources", config.PterodactylURL, config.ServerID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return "Error fetching server status"
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+config.APIToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return "Error fetching server status"
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		log.Printf("Error response from server: %s", body)
		return "Error fetching server status"
	}
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Printf("Error decoding JSON: %v", err)
		return "Error fetching server status"
	}
	status := result["attributes"].(map[string]interface{})["current_state"].(string)
	return fmt.Sprintf("Server status: %s", status)
}

// Note: This code assumes that the Pterodactyl API is set up correctly and that the server ID and API token are valid.
// Make sure to handle errors and edge cases in a production environment.
// Also, ensure that the config.yml file is properly formatted and contains the necessary fields.
