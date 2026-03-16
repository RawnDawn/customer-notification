package email

import (
	"log"
	"net/mail"
	"time"
)

// Connect to the SMTP to send promotional email
func SendPromotional(email *mail.Address, name string) error {

	log.Println("Sending email to:", email)

	time.Sleep(200 * time.Millisecond)

	log.Println("Email sent to:", email)

	return nil
}
