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
)

var logger = common.GetLogger()

func init() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	envir := env.GetEnv()
	router := core.New(envir)
	router.Add(func(group *gin.RouterGroup) {
		api.EmailRoute(group)
	})
	addBrokerSubscription(envir)

	err := router.Run(config.ServerConfiguration.Port)
	if err != nil {
		logger.Error(err, "server failed")
		os.Exit(1)
	}
}

func addBrokerSubscription(envir common.Env) {
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
