package client

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
	"io"
	"mri/client-email-sender/models"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func (cl *ClientEmailHandler) SendMessage(c *gin.Context) {
	var param models.ParamSendMessage
	_ = c.Bind(&param)
	dialer := cl.dialer()

	go func() {
		splitEmail := strings.Split(param.Recipients, ",")
		for _, email := range splitEmail {
			param.Recipients = email
			mailer := cl.mailer(param)
			err := dialer.DialAndSend(mailer)
			if err != nil {
				cl.Log.WriteToDbLog("gomail.v2", param.Subject, param.Recipients, param.Body, param.Attachment.Url, 500, fmt.Sprintf("Error: %v", err), fmt.Sprintf("Error: %v", err))
				fmt.Println(fmt.Sprintf("Error: %v", err.Error()))
			} else {
				cl.Log.WriteToDbLog("gomail.v2", param.Subject, param.Recipients, param.Body, param.Attachment.Url, 200, fmt.Sprintf("Error: %v", err), fmt.Sprintf("Error: %v", err))
			}
		}
	}()

	c.JSON(http.StatusOK, models.CommonResponse{
		Code:      200,
		IsSuccess: true,
		Message:   "Email still sending to " + param.Recipients,
	})
}

func downloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func (cl *ClientEmailHandler) dialer() *gomail.Dialer {
	port, _ := strconv.Atoi(cl.CONFIG_SMTP_PORT)
	dialer := gomail.NewDialer(
		cl.CONFIG_SMTP_HOST,
		port,
		cl.CONFIG_AUTH_EMAIL,
		cl.CONFIG_AUTH_PASSWORD,
	)
	return dialer
}

func (cl *ClientEmailHandler) mailer(param models.ParamSendMessage) *gomail.Message {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", cl.CONFIG_SENDER_NAME)
	mailer.SetHeader("To", param.Recipients)
	if param.RecipientsCC != "" {
		mailer.SetAddressHeader("Cc", param.RecipientsCC, param.Subject)
	}
	mailer.SetHeader("Subject", param.Subject)
	mailer.SetBody("text/html", param.Body)
	if param.Attachment.Url != "" {
		_ = downloadFile("./"+param.Attachment.Filename, param.Attachment.Url)
		mailer.Attach("./" + param.Attachment.Filename)
	}
	return mailer
}
