package email

type Email struct {
	To      string
	From    string
	Subject string
	Body    string
}

type EmailService interface {
	Send(email *Email) error
}
