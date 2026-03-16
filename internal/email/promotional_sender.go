package email

import (
	"log"
	"time"

	"github.com/rawndawn/customer-notification/internal/models"
)

// Connect to the SMTP to send promotional email
func SendPromotional(customer *models.Customer) error {
	// TODO - Email validation

	log.Println("Sending email to:", customer.Email)

	time.Sleep(200 * time.Millisecond)

	log.Println("Email sent to:", customer.Email)

	return nil
}
