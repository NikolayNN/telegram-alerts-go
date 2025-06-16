package alert

const Prefix = "[ALERT] "

// Msg adds the alert prefix to the message.
func Msg(msg string) string {
	return Prefix + msg
}
