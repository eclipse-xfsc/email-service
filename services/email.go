package services

import (
	"github.com/eclipse-xfsc/email-service/v2/common"
	"github.com/eclipse-xfsc/email-service/v2/config"
	"github.com/eclipse-xfsc/email-service/v2/model"
)

var logger = common.GetLogger()

func SendEmail(data *model.EmailData) error {

	logger.Logger.Info("preparing email send routine")

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
		logger.Logger.Error(nil, "emailSubject conformity error: "+data.MailSubject)
	}

	// check attachment parameters
	if len(data.EmailAttachmentBase64String) > 0 && len(data.EmailAttachmentName) < config.ServerConfiguration.Mail.SmtpAttachmentNameMin {
		logger.Logger.Error(nil, "emailAttachmentBase64String set, but no emailAttachmentName given")
		return nil
	}

	// check attachment parameters
	if len(data.EmailAttachmentBase64String) <= 0 && len(data.EmailAttachmentName) > 0 {
		logger.Logger.Error(nil, "emailAttachmentName given, but no emailAttachmentBase64String set")
		return nil
	}

	// TODO: implement meaningful checks

	return GetSMTP().SendEmail(data)

}
