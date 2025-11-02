package main

import (
	"log"
	"os"
	"path/filepath"
)

type Logger struct {
	Dir  string
	file *os.File
}

func (appLog *Logger) Initialize() error {

	if appLog.Dir == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			log.Fatalf("Unable to find user home directory: %v", err)
		}

		appLog.Dir = filepath.Join(home, "logs")

	}

	err := os.MkdirAll(appLog.Dir, 0755)
	if err != nil {
		log.Fatalf("Unable to create log directory: %v", err)
	}

	logPath := filepath.Join(appLog.Dir, LOG_FILE)

	appLog.file, err = os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	log.SetOutput(appLog.file)

	return nil
}

func (appLog *Logger) Close() {
	if appLog.file != nil {
		appLog.file.Close()
	}
}
