package common

import (
	log "github.com/sirupsen/logrus"
	"os"
)

var logger *log.Entry

type Logger struct {
}

func NewLogger() *log.Entry {
	log.SetFormatter(&log.JSONFormatter{})
	file, err := os.OpenFile("gapiparse.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Info("Failed to log to file, using default stderr")
	}

	logger = log.WithFields(log.Fields{
		"common": "this is app entry",
	})
	return logger
}
