package common

import (
	"log"
	"os"
	"sync"

	"github.com/eclipse-xfsc/email-service/config"
	core "github.com/eclipse-xfsc/microservice-core-go/pkg/logr"
)

var logger core.Logger

var once sync.Once

func GetLogger() core.Logger {
	once.Do(initLogger)
	return logger
}

func initLogger() {
	file, err := os.Create("log.txt") //should be replaced by something else
	l, err := core.New(config.ServerConfiguration.LogLevel, config.ServerConfiguration.IsDev, file)
	if err != nil {
		log.Fatal(err)
	}
	logger = *l
}
