package email

import (
	"errors"
	"log/slog"
	"os"
	"strconv"

	gomail "gopkg.in/gomail.v2"
)

type Email struct {
	From    string
	To      string
	Subject string
	Body    string
}

func NewEmail(
	to, subject, body string,
) *Email {
	return &Email{
		From:    os.Getenv("MAILER_EMAIL"),
		To:      to,
		Subject: subject,
		Body:    body,
	}
}

func (e *Email) Send() error {
	m := gomail.NewMessage()
	m.SetHeader("From", e.From)
	m.SetHeader("To", e.To)
	m.SetHeader("Subject", e.Subject)
	m.SetBody("text/html", e.Body)

	// Get port
	port, err := strconv.Atoi(os.Getenv("MAILER_PORT"))

	if err != nil {
		return errors.New("Cannot parse MAILER_PORT into int")
	}

	d := gomail.NewDialer(
		os.Getenv("MAILER_SERVER"),
		port,
		"",
		"",
	)

	// ssl, err := strconv.ParseBool(os.Getenv("MAILER_SSL"))
	// if err != nil {
	// 	return errors.New("Cannot parse MAILER_SSL into bool")
	// }

	// startTLS, err := strconv.ParseBool(os.Getenv("MAILER_START_TLS"))
	// if err != nil {
	// 	return errors.New("Cannot parse MAILER_START_TLS into bool")
	// }

	d.SSL = false
	d.TLSConfig = nil

	s, err := d.Dial()
	if err != nil {
		slog.Error("Cannot connect to mailer", slog.Any("err", err))
		return err
	}

	if err := gomail.Send(s, m); err != nil {
		slog.Error("Cannot send email", slog.Any("err", err))
		return err
	}

	return nil
}
