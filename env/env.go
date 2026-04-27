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

// EnvObj holds runtime environment state (brokers, swagger, etc.)
type EnvObj struct {
	brokers        map[string]*cloudeventprovider.CloudEventProviderClient
	swaggerOptions []func(config *ginSwagger.Config)
}

// NewEnv creates a fully initialized EnvObj (NO init(), no globals)
func NewEnv() *EnvObj {
	return &EnvObj{
		brokers: make(map[string]*cloudeventprovider.CloudEventProviderClient),
	}
}

// IsHealthy checks service readiness
func (env *EnvObj) IsHealthy() bool {
	return services.GetSMTP() != nil
}

// GetBroker returns a registered broker client
func (env *EnvObj) GetBroker(topic string) *cloudeventprovider.CloudEventProviderClient {
	if env.brokers == nil {
		return nil
	}
	return env.brokers[topic]
}

// BrokerSubscribe registers a NATS/CloudEvent subscription
func (env *EnvObj) BrokerSubscribe(topic string, handler func(e event.Event)) {
	if config.ServerConfiguration.Nats.URL == "" {
		logger.Error(nil, "NATS URL not configured")
		return
	}

	if env.brokers == nil {
		env.brokers = make(map[string]*cloudeventprovider.CloudEventProviderClient)
	}

	broker, err := connection.CloudEventsConnection(topic, cloudeventprovider.ConnectionTypeSub)
	if err != nil {
		logger.Error(err, "connection failed", "topic", topic)
		return
	}

	env.brokers[topic] = broker

	go func() {
		if err := broker.SubCtx(context.Background(), handler); err != nil {
			logger.Error(err, "subscription failed", "topic", topic)
		}
	}()
}

// SetSwaggerBasePath placeholder (kept for compatibility)
func (env *EnvObj) SetSwaggerBasePath(path string) {
	// intentionally empty for now
}

// SwaggerOptions returns configured swagger options
func (env *EnvObj) SwaggerOptions() []func(config *ginSwagger.Config) {
	return env.swaggerOptions
}