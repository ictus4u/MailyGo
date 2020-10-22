package main

import (
	"errors"
	"strings"

	"github.com/caarlos0/env/v6"
)

type config struct {
	Port              int      `env:"PORT" envDefault:"8080"`
	HoneyPots         []string `env:"HONEYPOTS" envDefault:"_t_email" envSeparator:","`
	DefaultRecipient  string   `env:"EMAIL_TO"`
	AllowedRecipients []string `env:"ALLOWED_TO" envSeparator:","`
	Sender            string   `env:"EMAIL_FROM"`
	SMTPUser          string   `env:"SMTP_USER"`
	SMTPPassword      string   `env:"SMTP_PASS"`
	SMTPHost          string   `env:"SMTP_HOST"`
	SMTPPort          int      `env:"SMTP_PORT" envDefault:"587"`
	GoogleAPIKey      string   `env:"GOOGLE_API_KEY"`
	Blacklist         string   `env:"BLACKLIST"  envDefault:"gambling,casino"`
	BlacklistArray    []string
}

func parseConfig() (*config, error) {
	cfg := &config{}
	if err := env.Parse(cfg); err != nil {
		return cfg, errors.New("failed to parse config")
	}
	cfg.BlacklistArray = strings.Split(cfg.Blacklist, ",")
	return cfg, nil
}

func checkRequiredConfig(cfg *config) bool {
	if cfg.DefaultRecipient == "" {
		return false
	}
	if len(cfg.AllowedRecipients) < 1 {
		return false
	}
	if cfg.Sender == "" {
		return false
	}
	if cfg.SMTPUser == "" {
		return false
	}
	if cfg.SMTPPassword == "" {
		return false
	}
	if cfg.SMTPHost == "" {
		return false
	}
	return true
}
