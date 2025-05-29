package interfaces

import (
	"net/http"

	"github.com/tantojos4/smtpapi-go/application"
	"github.com/tantojos4/smtpapi-go/domain/email"

	"github.com/labstack/echo/v4"
)

type EmailHandler struct {
	appService *application.EmailAppService
}

func NewEmailHandler(appService *application.EmailAppService) *EmailHandler {
	return &EmailHandler{appService: appService}
}

type EmailRequest struct {
	To      string `json:"to" validate:"required,email"`
	From    string `json:"from" validate:"required,email"`
	Subject string `json:"subject" validate:"required"`
	Body    string `json:"body" validate:"required"`
}

func (h *EmailHandler) SendEmail(c echo.Context) error {
	var req EmailRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request format"})
	}

	// Gunakan validator Echo yang sudah diregistrasi
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	email := &email.Email{
		To:      req.To,
		From:    req.From,
		Subject: req.Subject,
		Body:    req.Body,
	}

	if err := h.appService.SendEmail(email); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Email sent successfully"})
}
