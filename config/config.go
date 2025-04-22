package config

import (
	"github.com/kelseyhightower/envconfig"
)

var ServerConfiguration Config

func LoadConfig() error {
	conf := Config{}
	err := envconfig.Process("EMAIL", &conf)
	if err != nil {
		return err
	}
	ServerConfiguration = conf
	return nil
}

type Config struct {
	LogLevel string `default:"info"`
	IsDev    bool   `default:"true"`
	Name     string
	Tenant   string
	Port     int `default:"8080"`
	Nats     struct {
		Url          string
		QueueGroup   string
		TopicSend    string
		TopicReceive string
	}
	Mail struct {
		SmtpHost               string
		SmtpPort               string
		SmtpUsername           string
		SmtpPassword           string
		SmtpDefaultSubject     string
		SmtpDefaultSenderEmail string
		SmtpDefaultSenderName  string
		SmtpSubjectLengthMax   int `default:"255"`
		SmtpSubjectLengthMin   int `default:"5"`
		SmtpAttachmentNameMin  int `default:"5"`
	}
}
