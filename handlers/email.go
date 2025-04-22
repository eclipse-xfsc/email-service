package handlers

import (
	"context"
	"io"
	"net/http"

	cloudeventprovider "github.com/eclipse-xfsc/cloud-event-provider"
	"github.com/eclipse-xfsc/email-service/common"
	"github.com/eclipse-xfsc/email-service/config"
	"github.com/eclipse-xfsc/email-service/connection"
	"github.com/eclipse-xfsc/email-service/model"
	"github.com/eclipse-xfsc/email-service/services"
	"github.com/gin-gonic/gin"
)

const EmailSentEventType = "email.sent"

var logger = common.GetLogger()

func SendEmailNew(c *gin.Context) error {
	// incoming request, parse to new email send request
	logger.Debug("New Email Send Handler request")
	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		logger.Error(err, "email data invalid")
		return common.EmailError{Msg: "email data invalid", Code: http.StatusBadRequest}
	}

	eml, err := model.EmailDataFromJSONBytestream(data)

	if err != nil {
		logger.Error(err, "parsing email data failed")
		return common.EmailError{Msg: "email data invalid", Code: http.StatusBadRequest}
	}

	// send email
	err = SendEmail(eml)

	if err != nil {
		logger.Error(err, "could not send email")
		return common.EmailError{Msg: "could not send email", Code: http.StatusFailedDependency}
	}

	return nil
}

func SendEmailNewViaNats(c *gin.Context) error {
	// incoming request, parse to new email send request
	logger.Info("New Email Send via NATS Handler request")
	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		logger.Error(err, "email data invalid")
		return common.EmailError{Msg: "email data invalid", Code: http.StatusBadRequest}
	}

	_, err = model.EmailDataFromJSONBytestream(data)

	if err != nil {
		logger.Error(err, "email data invalid")
		return common.EmailError{Msg: "email data invalid", Code: http.StatusBadRequest}
	}

	e, err := cloudeventprovider.NewEvent(config.ServerConfiguration.Name, EmailSentEventType, data)
	if err != nil {
		logger.Error(err, "")
		return common.EmailError{Msg: "could not create event with provided email data", Code: http.StatusBadRequest}
	}
	// Publish the message
	broker, err := connection.CloudEventsConnection(config.ServerConfiguration.Nats.TopicSend, cloudeventprovider.ConnectionTypePub)
	if err != nil {
		return err
	}
	if err = broker.PubCtx(context.Background(), e); err != nil {
		logger.Error(err, "error publishing event")
		return common.EmailError{Msg: "error publishing event", Code: http.StatusFailedDependency}
	}
	logger.Debug("NATS message published", "topic", config.ServerConfiguration.Nats.TopicSend)

	return nil
}

func SendEmail(data *model.EmailData) error {

	logger.Info("preparing email send routine")

	// Check input values and compare with config values (defaults from config if given)
	emailSubjectDefault := config.ServerConfiguration.Mail.SmtpDefaultSubject
	if data.FromName == "" && emailSubjectDefault != "" {
		data.MailSubject = emailSubjectDefault
	}
	emailFromEmailaddressDefault := config.ServerConfiguration.Mail.SmtpDefaultSenderEmail
	if data.FromEmail == "" && emailFromEmailaddressDefault != "" {
		data.FromEmail = emailFromEmailaddressDefault
	}
	emailFromNameDefault := config.ServerConfiguration.Mail.SmtpDefaultSenderName
	if data.FromName == "" && emailFromNameDefault != "" {
		data.FromName = emailFromNameDefault
	}

	// check sanity of values and parameters
	if len(data.MailSubject) < config.ServerConfiguration.Mail.SmtpSubjectLengthMin || len(data.MailSubject) > config.ServerConfiguration.Mail.SmtpSubjectLengthMax {
		logger.Error(nil, "emailSubject conformity error: "+data.MailSubject)
	}

	// check attachment parameters
	if len(data.EmailAttachmentBase64String) > 0 && len(data.EmailAttachmentName) < config.ServerConfiguration.Mail.SmtpAttachmentNameMin {
		logger.Error(nil, "emailAttachmentBase64String set, but no emailAttachmentName given")
		return nil
	}

	// check attachment parameters
	if len(data.EmailAttachmentBase64String) <= 0 && len(data.EmailAttachmentName) > 0 {
		logger.Error(nil, "emailAttachmentName given, but no emailAttachmentBase64String set")
		return nil
	}

	// TODO: implement meaningful checks

	return services.GetSMTP().SendEmail(data)

}
