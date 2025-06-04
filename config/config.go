package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	DiscordToken   string `yaml:"discord_token"`
	CommandPrefix  string `yaml:"command_prefix"`
	APIToken       string `yaml:"api_token"`
	PterodactylURL string `yaml:"pterodactyl_url"`
	ServerID       string `yaml:"server_id"`
}

func LoadConfig() (*Config, error) {
	config := &Config{}
	configFile, err := os.ReadFile("config.yml")
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(configFile, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
