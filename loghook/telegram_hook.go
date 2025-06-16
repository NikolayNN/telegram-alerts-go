package loghook

import (
	"fmt"
	"telegram-alerts-go/alert"

	log "github.com/sirupsen/logrus"
	"telegram-alerts-go/telegram"
)

// TelegramHook перехватывает записи logrus и отправляет в Telegram,
// если marker == "ALERT".
type TelegramHook struct {
	Client     *telegram.Client
	MarkerKey  string
	ServiceTag string
}

func NewTelegramHook(client *telegram.Client, serviceTag string) *TelegramHook {
	return &TelegramHook{
		Client:     client,
		MarkerKey:  alert.ALERT_MARKER,
		ServiceTag: serviceTag,
	}
}

// Fire вызывается logrus при каждой записи.
func (h *TelegramHook) Fire(entry *log.Entry) error {
	val, ok := entry.Data[h.MarkerKey]
	if !ok || val != alert.ALERT_VALUE {
		return nil // нет нужного маркера
	}
	msg := fmt.Sprintf("[%s] %s", h.ServiceTag, entry.Message)
	return h.Client.SendMessage(msg)
}

// Levels указываем, что перехватываем все уровни.
func (h *TelegramHook) Levels() []log.Level {
	return log.AllLevels
}
