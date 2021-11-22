package main

import (
	"github.com/joho/godotenv"
	"mri/client-email-sender/handler"
	"mri/client-email-sender/handler/client"
)
func main() {
	_= godotenv.Load()
	server := handler.RunServer()

	handlerClient := client.NewClientEmailHandler()
	api := server.Group("api/v1/email")
	{
		api.POST("/send-notification-message", handlerClient.SendMessage)
	}

	server.Run(":8081")
}
