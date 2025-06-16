package config

import "os"

type Config struct {
	BotUsername string
	BotToken    string
	ChannelID   string
	ServiceName string
}

func LoadFromEnv() *Config {
	return &Config{
		BotUsername: os.Getenv("BOT_USERNAME"),
		BotToken:    os.Getenv("BOT_TOKEN"),
		ChannelID:   os.Getenv("CHANNEL_ID"),
		ServiceName: os.Getenv("SERVICE_NAME"),
	}
}
