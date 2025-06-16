package main

import (
	"telegram-alerts-go/alert"
	"telegram-alerts-go/config"
	"telegram-alerts-go/loghook"
	"telegram-alerts-go/telegram"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	cfg := config.LoadFromEnv()

	client := telegram.NewClient(cfg.BotToken, cfg.ChannelID)

	logger := zap.NewExample()
	logger = logger.WithOptions(zap.WrapCore(func(c zapcore.Core) zapcore.Core {
		return loghook.NewTelegramCore(c, client, cfg.ServiceName)
	}))

	logger.Info("Service started")
	logger.Error(alert.Msg("Test GO Telegram Alert"))
}
