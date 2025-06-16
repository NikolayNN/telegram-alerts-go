package main

import (
	log "github.com/sirupsen/logrus"
	"telegram-alerts-go/alert"
	"telegram-alerts-go/config"
	"telegram-alerts-go/loghook"
	"telegram-alerts-go/telegram"
)

func main() {

	cfg := config.LoadFromEnv()

	client := telegram.NewClient(cfg.BotToken, cfg.ChannelID)
	hook := loghook.NewTelegramHook(client, cfg.ServiceName)
	log.AddHook(hook)

	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})

	// Обычный лог
	log.Info("Service started")

	// АЛЕРТ-лог
	alert.Log().Error("Test GO Telegram Alert")
}
