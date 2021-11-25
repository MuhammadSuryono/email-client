package client

import (
	"github.com/gin-gonic/gin"
	"mri/client-email-sender/logs"
	"os"
)

type IClientEmail interface {
	SendMessage(c *gin.Context)
}

type ClientEmailHandler struct {
	CONFIG_SMTP_HOST     string
	CONFIG_SMTP_PORT     string
	CONFIG_SENDER_NAME   string
	CONFIG_AUTH_EMAIL    string
	CONFIG_AUTH_PASSWORD string
	Log                  logs.ILog
}

func NewClientEmailHandler() IClientEmail {
	log := logs.NewLog()
	return &ClientEmailHandler{
		CONFIG_SENDER_NAME:   os.Getenv("CONFIG_SENDER_NAME"),
		CONFIG_SMTP_HOST:     os.Getenv("CONFIG_SMTP_HOST"),
		CONFIG_AUTH_EMAIL:    os.Getenv("CONFIG_AUTH_EMAIL"),
		CONFIG_AUTH_PASSWORD: os.Getenv("CONFIG_AUTH_PASSWORD"),
		CONFIG_SMTP_PORT:     os.Getenv("CONFIG_SMTP_PORT"),
		Log:                  log,
	}
}
