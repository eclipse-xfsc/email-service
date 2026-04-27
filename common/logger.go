package common

import (
	"log"
	"os"
	"sync"

	"github.com/eclipse-xfsc/email-service/config"
	core "github.com/eclipse-xfsc/microservice-core-go/pkg/logr"
)

var (
	logger core.Logger
	once   sync.Once
)

// GetLogger returns a singleton logger instance
func GetLogger() core.Logger {
	once.Do(initLogger)
	return logger
}

// initLogger initializes the logger safely
func initLogger() {
	// Ensure config is already loaded before calling this function
	// (important: config.LoadConfig() must happen in main before GetLogger)

	logLevel := config.ServerConfiguration.LogLevel
	isDev := config.ServerConfiguration.IsDev

	// IMPORTANT: use stdout for Docker/Kubernetes logging
	output := os.Stdout

	l, err := core.New(logLevel, isDev, output)
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}

	logger = *l

	logger.Info("logger initialized successfully")
}