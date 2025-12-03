package common

import (
	"github.com/eclipse-xfsc/email-service/v2/config"
	core "github.com/eclipse-xfsc/microservice-core-go/pkg/logr"
)

var logger core.Logger

func GetLogger() core.Logger {
	return logger
}

func InitLogger() {
	log, err := core.New(config.ServerConfiguration.LogLevel, config.ServerConfiguration.IsDev, nil)
	if err != nil {
		panic(err)
	}

	logger = *log
}
