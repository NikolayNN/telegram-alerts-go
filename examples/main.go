package main

import (
	"telegram-alerts-go/alert"
	"telegram-alerts-go/config"
	"telegram-alerts-go/loghook"

	"go.uber.org/zap"
)

func main() {
	cfg := config.LoadFromEnv()

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger = loghook.AttachToLogger(logger, cfg)

	// Обычный лог
	logger.Info("Service started")

	// АЛЕРТ-лог
	logger.Error(alert.Prefix("Test GO Telegram Alert"))
}
