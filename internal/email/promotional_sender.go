package email

import (
	"errors"
	"log"
	"time"
)

var ErrInvalidCustomerEmail = errors.New("Invalid customer email")

// Connect to the SMTP to send promotional email
func SendPromotional(email, name string) error {

	log.Println("Sending email to:", email)

	time.Sleep(200 * time.Millisecond)

	log.Println("Email sent to:", email)

	return nil
}
