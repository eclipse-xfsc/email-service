package env

import (
	"context"

	"github.com/cloudevents/sdk-go/v2/event"
	cloudeventprovider "github.com/eclipse-xfsc/cloud-event-provider"
	"github.com/eclipse-xfsc/email-service/common"
	"github.com/eclipse-xfsc/email-service/config"
	"github.com/eclipse-xfsc/email-service/connection"
	"github.com/eclipse-xfsc/email-service/services"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var logger = common.GetLogger()

func init() {
	env = &EnvObj{}
}

func GetEnv() *EnvObj {
	return env
}

var env *EnvObj

type EnvObj struct {
	brokers        map[string]*cloudeventprovider.CloudEventProviderClient
	swaggerOptions []func(config *ginSwagger.Config)
}

func (env *EnvObj) IsHealthy() bool {
	return services.GetSMTP() != nil
}

func (env *EnvObj) GetBroker(topic string) *cloudeventprovider.CloudEventProviderClient {
	return env.brokers[topic]
}
func (env *EnvObj) BrokerSubscribe(topic string, handler func(e event.Event)) {
	if config.ServerConfiguration.Nats.Url != "" {
		broker, err := connection.CloudEventsConnection(topic, cloudeventprovider.ConnectionTypeSub)
		if err != nil {
			logger.Error(err, "connection failed", "topic", topic)
			return
		}
		env.brokers[topic] = broker
		go func() {
			er := broker.SubCtx(context.Background(), handler)
			if er != nil {
				logger.Error(er, "subscription failed", "topic", topic)
			}
		}()
	}
}

func (env *EnvObj) SetSwaggerBasePath(path string) {

}

// SwaggerOptions swagger config options. See https://github.com/swaggo/gin-swagger?tab=readme-ov-file#configuration
func (env *EnvObj) SwaggerOptions() []func(config *ginSwagger.Config) {
	return env.swaggerOptions
}
