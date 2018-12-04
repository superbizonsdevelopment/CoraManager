package manager

import (
	"net/smtp"

	"../basic"
)

func SendMail(user *basic.User) error {
	auth := smtp.PlainAuth("", "@example.com", "password", "mail.example.com")

	to := []string{user.Email}

	msg := []byte("To: " + user.Email + "\r\n" + "Subject: discount Gophers!\r\n" + "\r\n" + "This is the email body.\r\n")

	err := smtp.SendMail("mail.example.com:25", auth, "sender@example.org", to, msg)

	if err != nil {
		return err
	}

	return nil
}
