package main

import (
	"log"
	"os"

	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/eclipse-xfsc/email-service/api"
	"github.com/eclipse-xfsc/email-service/common"
	"github.com/eclipse-xfsc/email-service/config"
	"github.com/eclipse-xfsc/email-service/env"
	"github.com/eclipse-xfsc/email-service/handlers"
	"github.com/eclipse-xfsc/email-service/model"
	core "github.com/eclipse-xfsc/microservice-core-go/pkg/server"
	"github.com/gin-gonic/gin"
	logr "github.com/eclipse-xfsc/microservice-core-go/pkg/logr" 
	services "github.com/eclipse-xfsc/email-service/services"
)


func main() {
	if err := config.LoadConfig(); err != nil {
		log.Fatal(err)
	}

	logger := common.GetLogger()
	logger.Info("starting service")

	if err := services.InitSMTP(); err != nil {
		logger.Error(err, "could not initialise SMTP server")
		os.Exit(1)
	}

	envir := env.NewEnv()

	router := core.New(envir)

	router.Add(func(group *gin.RouterGroup) {
		api.EmailRoute(group)
	})

	addBrokerSubscription(envir, logger)

	if err := router.Run(config.ServerConfiguration.App.Port); err != nil {
		logger.Error(err, "server failed")
		os.Exit(1)
	}
}

func addBrokerSubscription(envir common.Env, logger logr.Logger) {
	envir.BrokerSubscribe(config.ServerConfiguration.Nats.TopicReceive, func(e event.Event) {
		var eml model.EmailData
		err := e.DataAs(&eml)
		if err != nil {
			logger.Error(err, "error during handling cloudevents subscription")
		}
		err = handlers.SendEmail(&eml)
		if err != nil {
			logger.Error(err, "error during handling cloudevents subscription")
		}
	})
}
