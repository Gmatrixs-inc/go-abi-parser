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

	// Output to stdout instead of the default stderr
	// You could set this to any `io.Writer` such as a file
	file, err := os.OpenFile("gapiparse.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
	// Can be any io.Writer, see below for File example

	// Only log the warning severity or above.
	//log.SetLevel(log.WarnLevel)
	logger = log.WithFields(log.Fields{
		"common": "this is app entry",
	})
	return logger
}
