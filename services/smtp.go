package services

import (
	"github.com/eclipse-xfsc/email-service/common"
	"github.com/eclipse-xfsc/email-service/config"
	"github.com/eclipse-xfsc/email-service/model"
	"encoding/base64"
	log "github.com/sirupsen/logrus"
	mail "github.com/xhit/go-simple-mail/v2"
	"os"
	"strconv"
)

func GetSMTP() *SMTP {
	return smtp
}

var smtp *SMTP

func init() {
	var err error
	smtp, err = newSMTP()
	if err != nil {
		var logger = common.GetLogger()
		logger.Error(err, "could not initialise SMTP server")
		os.Exit(1)
	}
}

type SMTP struct {
	client *mail.SMTPClient
}

func (s *SMTP) SendEmail(data *model.EmailData) error {
	// Create Email
	email := mail.NewMSG()
	email.SetFrom("" + data.FromName + "<" + data.FromEmail + ">")
	email.AddTo("" + data.ToName + "<" + data.ToEmail + ">")
	email.SetSubject(data.MailSubject)
	email.SetBody(mail.TextHTML, data.MailBody)

	if len(data.EmailAttachmentBase64String) > 0 {
		attachment, _ := base64.StdEncoding.DecodeString(data.EmailAttachmentBase64String)
		email.Attach(&mail.File{Data: attachment, Name: data.EmailAttachmentName, Inline: true})
		log.Infof("Attachment added to Email: %s", data.EmailAttachmentName)
	} else {
		log.Info("Empty attachment definition found - no attachment added to email")
	}

	// Send Email
	log.Debug("sending email")
	err := email.Send(s.client)
	if err != nil {
		log.Error(err)
	} else {
		log.Debug("email sent")
	}

	return err
}
func newSMTP() (*SMTP, error) {
	// Apply Configuration Values
	server := mail.NewSMTPClient()
	server.Host = config.ServerConfiguration.Mail.SmtpHost
	server.Port, _ = strconv.Atoi(config.ServerConfiguration.Mail.SmtpPort)
	server.Username = config.ServerConfiguration.Mail.SmtpUsername
	server.Password = config.ServerConfiguration.Mail.SmtpPassword
	server.Encryption = mail.EncryptionNone // mail.EncryptionTLS

	// Initialize SMTP Client
	log.Debug("connecting to SMTP server")
	smtpClient, err := server.Connect()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &SMTP{client: smtpClient}, nil
}
