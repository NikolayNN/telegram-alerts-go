package loghook

import (
	"fmt"
	"strings"

	"telegram-alerts-go/alert"
	"telegram-alerts-go/telegram"

	"go.uber.org/zap/zapcore"
)

type TelegramCore struct {
	core       zapcore.Core
	client     *telegram.Client
	serviceTag string
	prefix     string
}

func NewTelegramCore(base zapcore.Core, client *telegram.Client, serviceTag string) zapcore.Core {
	return &TelegramCore{
		core:       base,
		client:     client,
		serviceTag: serviceTag,
		prefix:     alert.Prefix,
	}
}

func (t *TelegramCore) Enabled(lvl zapcore.Level) bool {
	return t.core.Enabled(lvl)
}

func (t *TelegramCore) With(fields []zapcore.Field) zapcore.Core {
	return &TelegramCore{
		core:       t.core.With(fields),
		client:     t.client,
		serviceTag: t.serviceTag,
		prefix:     t.prefix,
	}
}

func (t *TelegramCore) Check(ent zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if t.Enabled(ent.Level) {
		return t.core.Check(ent, ce)
	}
	return ce
}

func (t *TelegramCore) Write(ent zapcore.Entry, fields []zapcore.Field) error {
	msg := ent.Message
	if strings.HasPrefix(msg, t.prefix) {
		trimmed := strings.TrimPrefix(msg, t.prefix)
		ent.Message = trimmed
		var emoji string
		switch ent.Level {
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
		text := fmt.Sprintf("%s\n [%s]- %s", t.serviceTag, emoji, trimmed)
		_ = t.client.SendMessage(text)
	}
	return t.core.Write(ent, fields)
}

func (t *TelegramCore) Sync() error {
	return t.core.Sync()
}
