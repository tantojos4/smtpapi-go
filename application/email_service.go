package application

import (
	"github.com/tantojos4/smtpapi-go/domain/email"
)

type EmailAppService struct {
	emailService email.EmailService
}

func NewEmailAppService(emailService email.EmailService) *EmailAppService {
	return &EmailAppService{emailService: emailService}
}

func (s *EmailAppService) SendEmail(e *email.Email) error {
	// Add business logic here if needed
	return s.emailService.Send(e)
}
