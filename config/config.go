package config

import (
	"errors"
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

var ServerConfiguration Config

// LoadConfig loads and validates environment configuration
func LoadConfig() error {
	var cfg Config

	// parse environment variables
	if err := envconfig.Process("", &cfg); err != nil {
		return fmt.Errorf("failed to load env: %w", err)
	}

	// validate config
	if err := validate(cfg); err != nil {
		return fmt.Errorf("invalid config: %w", err)
	}

	ServerConfiguration = cfg
	return nil
}

type Config struct {
	// App-level config (single source of truth)
	App struct {
		Name string `envconfig:"APP_NAME" required:"true"`
		Env  string `envconfig:"ENV" default:"dev"`
		Port int    `envconfig:"PORT" default:"8080"`
	}

	LogLevel string `envconfig:"LOG_LEVEL" default:"info"`
	IsDev    bool   `envconfig:"IS_DEV" default:"true"`

	Nats NatsConfig
	SMTP SMTPConfig
}

type NatsConfig struct {
	URL          string `envconfig:"NATS_URL" required:"true"`
	QueueGroup   string `envconfig:"NATS_QUEUE_GROUP" required:"true"`
	TopicSend    string `envconfig:"NATS_TOPIC_SEND" required:"true"`
	TopicReceive string `envconfig:"NATS_TOPIC_RECEIVE" required:"true"`
}

type SMTPConfig struct {
	Host     string `envconfig:"SMTP_HOST" required:"true"`
	Port     int    `envconfig:"SMTP_PORT" required:"true"`

	Username string `envconfig:"SMTP_USERNAME"`
	Password string `envconfig:"SMTP_PASSWORD"`

	DefaultSenderEmail string `envconfig:"SMTP_DEFAULT_SENDER_EMAIL" required:"true"`
	DefaultSenderName  string `envconfig:"SMTP_DEFAULT_SENDER_NAME"`
	DefaultSubject     string `envconfig:"SMTP_DEFAULT_SUBJECT" default:"No Subject"`

	SubjectMin int `envconfig:"SMTP_SUBJECT_LENGTH_MIN" default:"5"`
	SubjectMax int `envconfig:"SMTP_SUBJECT_LENGTH_MAX" default:"255"`
}

func validate(c Config) error {
	if c.App.Name == "" {
		return errors.New("APP_NAME is required")
	}

	if c.App.Port <= 0 || c.App.Port > 65535 {
		return errors.New("PORT must be between 1-65535")
	}

	if c.Nats.URL == "" {
		return errors.New("NATS_URL is required")
	}

	if c.Nats.QueueGroup == "" {
		return errors.New("NATS_QUEUE_GROUP is required")
	}

	if c.SMTP.Host == "" {
		return errors.New("SMTP_HOST is required")
	}

	if c.SMTP.Port <= 0 {
		return errors.New("SMTP_PORT must be > 0")
	}

	if c.SMTP.DefaultSenderEmail == "" {
		return errors.New("SMTP_DEFAULT_SENDER_EMAIL is required")
	}

	if c.SMTP.SubjectMin <= 0 || c.SMTP.SubjectMax <= 0 {
		return errors.New("SMTP subject length config invalid")
	}

	if c.SMTP.SubjectMin > c.SMTP.SubjectMax {
		return errors.New("SMTP subject min cannot be greater than max")
	}

	return nil
}