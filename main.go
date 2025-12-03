package main

import (
	"log"

	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/eclipse-xfsc/email-service/v2/common"
	"github.com/eclipse-xfsc/email-service/v2/config"
	"github.com/eclipse-xfsc/email-service/v2/env"
	"github.com/eclipse-xfsc/email-service/v2/pkg/messaging"
	"github.com/eclipse-xfsc/email-service/v2/services"
)

var logger = common.GetLogger()

func init() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	common.InitLogger()
	envir := env.GetEnv()
	err := services.CreateSmtp()
	if err != nil {
		log.Fatal(err)
	}

	addBrokerSubscription(envir)
}

func addBrokerSubscription(envir common.Env) {
	envir.BrokerSubscribe(config.ServerConfiguration.Nats.TopicReceive, func(e event.Event) {
		var eml messaging.EmailPublication
		err := e.DataAs(&eml)
		if err != nil {
			logger.Logger.Error(err, "error during handling cloudevents subscription")
		}
		err = services.SendEmail(&eml.Email)
		if err != nil {
			logger.Logger.Error(err, "error during handling cloudevents subscription")
		}
	})
}
