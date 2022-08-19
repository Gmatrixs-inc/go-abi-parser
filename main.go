package main

import (
	"gabiparser/cmd"
	log "github.com/sirupsen/logrus"
	"os"
)

var logger *log.Entry

func init() {

	// Log as JSON instead of the default ASCII formatter.
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
}

func main() {
	logger.Info("Start Processing...")
	err := cmd.Execute()
	if err != nil {
		logger.Error("Failed to process data")
		return
	}
}
