package logger

import (
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	once   sync.Once
	logger *logrus.Logger
)

func init() {
	once.Do(func() {
		logrus.SetFormatter(&logrus.JSONFormatter{})
		logrus.SetOutput(os.Stdout)
		logrus.SetLevel(logrus.DebugLevel)

		logger = logrus.New()
	})
}

func Logger() *logrus.Logger {
	return logger
}
