# Telegram Alerts for Go

This project provides a minimal helper to deliver alert logs from a Go application to a Telegram chat. It exposes:

- **telegram** client for sending messages using the Telegram Bot API
- **loghook** package with a zap hook that forwards prefixed entries
- **alert** helper that prefixes alert messages with `[ALERT]`
- **config** loader that reads required settings from environment variables
- Alert messages are prefixed with an emoji depending on the log level
  (💚 for INFO, 💛 for WARN and 💔 for ERROR)

## Quick start

1. Create a Telegram bot and obtain its token
2. Create a chat (or a channel) and get its numeric ID
3. Copy `.env.example` to `.env` and fill in your values

```bash
cp .env.example .env
```

4. Add the Telegram hook to your zap logger:

```go
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

    // Regular log message
    logger.Info("Service started")

    // Alert log message
logger.Error(alert.Prefix("Database is down!"))
}
```

`AttachToLogger` checks that the required environment variables are set. If any
are missing, it logs a warning and disables Telegram forwarding.

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

