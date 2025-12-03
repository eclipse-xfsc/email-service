package messaging

import (
	"github.com/eclipse-xfsc/email-service/v2/model"
	"github.com/eclipse-xfsc/nats-message-library/common"
)

const EmailSentEventType = "email.sent"
const EmailSentEventTopic = "email"

type EmailPublication struct {
	common.Request
	Email model.EmailData
}
