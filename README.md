# telegramalert

Лёгкая библиотека для отправки логов-алертов из Go-приложений в Telegram.

```
// Загружаем настройки из переменных окружения
cfg := config.LoadFromEnv()

// Регистрируем Telegram-хуку
client := telegram.NewClient(cfg.BotToken, cfg.ChannelID)
hook   := loghook.NewTelegramHook(client, "marker", cfg.ServiceName)
log.AddHook(hook)

// Лог без маркера
log.Info("Сервис запущен")

// Лог-алерт
alertlog.Alert().Error("Падает база данных!")
```

### 3. `.env.example`

```dotenv
BOT_TOKEN=123456789:AAAbbbCCCdddEEEfffGGG
CHANNEL_ID=-1001234567890
SERVICE_NAME=service_name
