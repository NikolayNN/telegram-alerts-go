package config

import "os"

type Config struct {
	BotUsername string
	BotToken    string
	ChannelID   string
	ServiceName string
}

// MissingFields returns names of the environment variables that were not set.
func (c *Config) MissingFields() []string {
	missing := []string{}
	if c.BotToken == "" {
		missing = append(missing, "BOT_TOKEN")
	}
	if c.ChannelID == "" {
		missing = append(missing, "CHANNEL_ID")
	}
	if c.ServiceName == "" {
		missing = append(missing, "SERVICE_NAME")
	}
	if c.BotUsername == "" {
		missing = append(missing, "BOT_USERNAME")
	}
	return missing
}

func LoadFromEnv() *Config {
	return &Config{
		BotUsername: os.Getenv("BOT_USERNAME"),
		BotToken:    os.Getenv("BOT_TOKEN"),
		ChannelID:   os.Getenv("CHANNEL_ID"),
		ServiceName: os.Getenv("SERVICE_NAME"),
	}
}
