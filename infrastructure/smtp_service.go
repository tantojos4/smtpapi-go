package infrastructure

import (
	"fmt"
	"net/smtp"

	"github.com/tantojos4/smtpapi-go/domain/email"
)

type SMTPConfig struct {
	Host     string
	Port     string
	Username string
	Password string
}

type SMTPEmailService struct {
	config *SMTPConfig
}

func NewSMTPEmailService(config *SMTPConfig) email.EmailService {
	return &SMTPEmailService{config: config}
}

func (s *SMTPEmailService) Send(e *email.Email) error {
	if s.config.Username == "" || s.config.Password == "" {
		return fmt.Errorf("SMTP credentials not configured")
	}

	auth := smtp.PlainAuth("", s.config.Username, s.config.Password, s.config.Host)
	to := []string{e.To}
	msg := []byte(fmt.Sprintf(
		"To: %s\r\n"+
			"Subject: %s\r\n"+
			"\r\n"+
			"%s\r\n", e.To, e.Subject, e.Body))

	return smtp.SendMail(
		s.config.Host+":"+s.config.Port,
		auth,
		e.From,
		to,
		msg)
}
