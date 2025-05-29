package interfaces

import (
	"net/http"

	"github.com/tantojos4/smtpapi-go/application"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// CustomValidator untuk registrasi validator
type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func NewRouter(appService *application.EmailAppService) *echo.Echo {
	e := echo.New()

	// Setup validator
	e.Validator = &CustomValidator{validator: validator.New()}

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		if he, ok := err.(*echo.HTTPError); ok {
			// Custom error handling
			c.JSON(he.Code, map[string]interface{}{
				"error":  he.Message.(string),
				"detail": he.Internal,
			})
		} else if ve, ok := err.(validator.ValidationErrors); ok {
			errors := make(map[string]string)
			for _, fe := range ve {
				errors[fe.Field()] = msgForTag(fe.Tag())
			}
			c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error":   "Validation failed",
				"details": errors,
			})
		} else {
			// Default error handling
			c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error":  "Internal Server Error",
				"detail": err.Error(),
			})
		}
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	handler := NewEmailHandler(appService)

	e.POST("/send-email", handler.SendEmail)

	return e
}
func msgForTag(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	}
	return ""
}
