package main

import (
	"github.com/tantojos4/smtpapi-go/application"
	"github.com/tantojos4/smtpapi-go/infrastructure"
	"github.com/tantojos4/smtpapi-go/interfaces"
)

func main() {
	smtpConfig := infrastructure.LoadSMTPConfig()
	emailService := infrastructure.NewSMTPEmailService(smtpConfig)
	appService := application.NewEmailAppService(emailService)

	e := interfaces.NewRouter(appService)

	serverPort := ":" + infrastructure.GetEnv("SERVER_PORT", "8080")
	e.Logger.Fatal(e.Start(serverPort))
}
