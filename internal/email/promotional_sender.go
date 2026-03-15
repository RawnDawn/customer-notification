package email

import (
	"log"
	"time"
)

type PromotionalSender struct{}

func NewPromotionalSender() *PromotionalSender {
	return &PromotionalSender{}
}

// Connect to the SMTP to send promotional email
func (f *PromotionalSender) Send(to, subject, body string) error {

	log.Println("Sending email to:", to)

	time.Sleep(200 * time.Millisecond)

	log.Println("Email sent to:", to)

	return nil
}
