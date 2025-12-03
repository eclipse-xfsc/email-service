package connection

import (
	"fmt"

	cloudeventprovider "github.com/eclipse-xfsc/cloud-event-provider"
	"github.com/eclipse-xfsc/email-service/v2/common"
	"github.com/eclipse-xfsc/email-service/v2/config"
)

var logger = common.GetLogger()

func CloudEventsConnection(topic string, typ cloudeventprovider.ConnectionType) (*cloudeventprovider.CloudEventProviderClient, error) {
	client, err := cloudeventprovider.New(cloudeventprovider.Config{
		Protocol: cloudeventprovider.ProtocolTypeNats,
		Settings: cloudeventprovider.NatsConfig{
			Url:        config.ServerConfiguration.Nats.Url,
			QueueGroup: config.ServerConfiguration.Nats.QueueGroup,
		},
	}, typ, topic)

	if err != nil {
		logger.Logger.Error(err, "error during establishing cloudevents connection")
		return nil, err
	} else {
		logger.Logger.Info(fmt.Sprintf("cloudEvents can be received over topic: %s", topic))
	}
	return client, nil
}
