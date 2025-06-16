package loghook

import (
	"fmt"
	"strings"

	"telegram-alerts-go/alert"
	"telegram-alerts-go/config"
	"telegram-alerts-go/telegram"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewTelegramHook returns a zap.Option that forwards log entries with
// alert prefix to Telegram.
func NewTelegramHook(client *telegram.Client, serviceTag string) zap.Option {
	return zap.Hooks(func(entry zapcore.Entry) error {
		if !strings.HasPrefix(entry.Message, alert.ALERT_PREFIX) {
			return nil
		}
		msg := strings.TrimPrefix(entry.Message, alert.ALERT_PREFIX)
		msg = strings.TrimSpace(msg)

		var emoji string
		switch entry.Level {
		case zapcore.InfoLevel:
			emoji = "\xF0\x9F\x92\x9A" // 💚
		case zapcore.WarnLevel:
			emoji = "\xF0\x9F\x92\x9B" // 💛
		case zapcore.ErrorLevel, zapcore.DPanicLevel, zapcore.PanicLevel, zapcore.FatalLevel:
			emoji = "\xF0\x9F\x92\x94" // 💔
		}
		if emoji != "" {
			emoji += " "
		}
		telegramMsg := fmt.Sprintf("%s\n [%s] - %s", serviceTag, emoji, msg)
		return client.SendMessage(telegramMsg)
	})
}

// AttachToLogger checks the configuration and attaches the Telegram hook if all
// required fields are present. Otherwise it logs a warning listing missing
// environment variables.
func AttachToLogger(logger *zap.Logger, cfg *config.Config) *zap.Logger {
	missing := cfg.MissingFields()
	if len(missing) == 0 {
		client := telegram.NewClient(cfg.BotToken, cfg.ChannelID)
		return logger.WithOptions(NewTelegramHook(client, cfg.ServiceName))
	}
	logger.Warn("telegram is not configured", zap.Strings("missing_env", missing))
	return logger
}
