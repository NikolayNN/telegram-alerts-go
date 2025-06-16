package main

import (
	"telegram-alerts-go/alert"
	"telegram-alerts-go/config"
	"telegram-alerts-go/loghook"
	"telegram-alerts-go/telegram"

	"go.uber.org/zap"
)

func main() {
	cfg := config.LoadFromEnv()

	client := telegram.NewClient(cfg.BotToken, cfg.ChannelID)

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger = logger.WithOptions(loghook.NewTelegramHook(client, cfg.ServiceName))

	// Обычный лог
	logger.Info("Service started")

	// АЛЕРТ-лог
	logger.Error(alert.Prefix("Test GO Telegram Alert"))
}
