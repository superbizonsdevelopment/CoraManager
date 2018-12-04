package manager

import (
	"fmt"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	api "github.com/superbizonsdevelopment/CoraManager/api"
)

func SendMail(user *api.User) error {
	configuration, err := LoadConfiguration()

	from := mail.NewEmail("SuperbizonsDevelopment", configuration.Mail.Username)
	subject := "TestTopic"
	to := mail.NewEmail(user.Name, user.Email)
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		return err
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
	return nil
}
