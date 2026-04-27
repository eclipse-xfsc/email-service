package services

import (
	"github.com/eclipse-xfsc/email-service/config"
	"github.com/eclipse-xfsc/email-service/model"
	"encoding/base64"
	log "github.com/sirupsen/logrus"
	mail "github.com/xhit/go-simple-mail/v2"
	"fmt"
)

func GetSMTP() *SMTP {
	return smtp
}

var smtp *SMTP

func InitSMTP() error {
	var err error
	smtp, err = newSMTP()
	if err != nil {
		return err
	}
	return nil
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
	cfg := config.ServerConfiguration.SMTP

	if cfg.Host == "" || cfg.Port == 0 {
		return nil, fmt.Errorf("invalid SMTP config: host=%q port=%d", cfg.Host, cfg.Port)
	}
	// Apply Configuration Values
	server := mail.NewSMTPClient()
	server.Host = config.ServerConfiguration.SMTP.Host
	server.Port = config.ServerConfiguration.SMTP.Port
	server.Username = config.ServerConfiguration.SMTP.Username
	server.Password = config.ServerConfiguration.SMTP.Password
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
