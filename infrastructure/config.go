package infrastructure

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Using system environment variables")
	}
}

// type SMTPConfig struct {
//     Host     string
//     Port     string
//     Username string
//     Password string
// }

// Fungsi GetEnv yang bisa diakses publik
func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func LoadSMTPConfig() *SMTPConfig {
	return &SMTPConfig{
		Host:     GetEnv("SMTP_HOST", "smtp.gmail.com"),
		Port:     GetEnv("SMTP_PORT", "587"),
		Username: GetEnv("SMTP_USERNAME", ""),
		Password: GetEnv("SMTP_PASSWORD", ""),
	}
}
