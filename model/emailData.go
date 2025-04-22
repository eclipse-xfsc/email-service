package model

import (
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
)

type EmailData struct {
	MailSubject                 string `json:"subject"`
	FromName                    string `json:"from-name"`
	FromEmail                   string `json:"from-email"`
	ToName                      string `json:"to-name"`
	ToEmail                     string `json:"to-email"`
	MailBody                    string `json:"body"`
	EmailAttachmentName         string `json:"attachment-name" default:""`
	EmailAttachmentBase64String string `json:"attachment-base64" default:""`
}

func EmailDataFromJSONBytestream(data []byte) (*EmailData, error) {
	log.Debug("parsing EmailData from JSON raw bytestream: " + string(data))

	var eml EmailData
	// use json decoder that throws an error if the data does not match:

	err := json.Unmarshal(data, &eml)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	log.Debugf("unmarshaled json - subject: %s, fromName: %s, fromEmail: %s, toName: %s, toEmail: %s, body: %s, attachmentName: %s, attachmentBase64; %s",
		eml.MailSubject, eml.FromName, eml.FromEmail, eml.ToName, eml.ToEmail, eml.MailBody, eml.EmailAttachmentName, eml.EmailAttachmentBase64String)

	if err == nil {
		log.Debug(eml)

		// make sure mandatory field subject is set
		if eml.MailSubject == "" {
			log.Error("Mandatory field subject missing in JSON structure")
			return nil, fmt.Errorf("mandatory field mailSubject missing in json structure")
		}

		// make sure mandatory field toEmail is set
		if eml.ToEmail == "" {
			log.Error("Mandatory field to-email missing in JSON structure")
			return nil, fmt.Errorf("mandatory field toEmail missing in json structure")
		}

		// make sure mandatory field body is set
		if eml.MailBody == "" {
			log.Error("Mandatory field body missing in JSON structure")
			return nil, fmt.Errorf("mandatory field body missing in json structure")
		}
	}

	return &eml, err
}
