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

	from := mail.NewEmail(configuration.Username, configuration.Email)
	subject := "I <3 GoLang"
	to := mail.NewEmail(user.Name, user.Email)
	plainTextContent := "Golang The Best Programming Language"
	htmlContent := "<strong>Test</strong>"
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
