package notification

import (
	"log"
	"net/mail"

	"github.com/rawndawn/customer-notification/internal/infrestructure/mailer"
)

// Connect to the SMTP to send promotional email
func SendPromotional(email *mail.Address, name string) error {

	log.Println("Sending email to:", email)

	err := mailer.NewEmail(
		email.Address,
		"Promotional",
		"Hi "+name+", welcome to our service!",
	).
		Send()

	if err != nil {
		return err
	}

	return nil
}
