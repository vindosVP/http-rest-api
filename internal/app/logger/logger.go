package logger

import (
	"github.com/sirupsen/logrus"
)

var (
	logger *logrus.Logger
)

func GetLogger() *logrus.Logger {
	return logger
}

func ConfigureLogger(logLvl string) error {
	if logger == nil {
		logger = logrus.New()
	}

	lev, err := logrus.ParseLevel(logLvl)

	if err != nil {
		return err
	}

	logger.SetLevel(lev)
	return nil
}
