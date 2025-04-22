package common

import "github.com/cloudevents/sdk-go/v2/event"

type Env interface {
	IsHealthy() bool
	BrokerSubscribe(string, func(e event.Event))
}
