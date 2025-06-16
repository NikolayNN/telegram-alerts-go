package alert

import logger "github.com/sirupsen/logrus"

const ALERT_MARKER = "ALERT_MARKER"
const ALERT_VALUE = "TELEGRAM"

func Log() *logger.Entry {
	return logger.WithField(ALERT_MARKER, ALERT_VALUE)
}
