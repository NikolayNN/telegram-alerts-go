package alert

// ALERT_PREFIX is added to log messages that should be forwarded to Telegram.
const ALERT_PREFIX = "[ALERT]"

// Prefix prepends ALERT_PREFIX to the provided message.
func Prefix(msg string) string {
	return ALERT_PREFIX + " " + msg
}
