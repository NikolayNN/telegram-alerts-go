# Telegram Alerts for Go

This project provides a minimal helper to deliver alert logs from a Go application to a Telegram chat. It exposes:

- **telegram** client for sending messages using the Telegram Bot API
- **loghook** package with a logrus hook that forwards marked entries
- **alert** helper that marks log records as alerts
- **config** loader that reads required settings from environment variables

## Quick start

1. Create a Telegram bot and obtain its token
2. Create a chat (or a channel) and get its numeric ID
3. Copy `.env.example` to `.env` and fill in your values

```bash
cp .env.example .env
```

4. Add the Telegram hook to your logrus logger:

```go
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

    // Regular log message
    log.Info("Service started")

    // Alert log message
    alert.Log().Error("Database is down!")
}
```

## Environment variables

| Name           | Description                         |
|----------------|-------------------------------------|
| `BOT_TOKEN`    | Bot token received from @BotFather  |
| `BOT_USERNAME` | Optional bot username               |
| `CHANNEL_ID`   | Chat or channel ID where alerts go  |
| `SERVICE_NAME` | Identifier added to alert messages  |

See `.env.example` for a sample configuration file.

## Example

A ready-to-run example is available in [`examples/main.go`](examples/main.go).
Run it after setting the environment variables to test the alerting pipeline.

