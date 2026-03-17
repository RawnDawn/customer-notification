package email

import (
	"crypto/tls"
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
		From:    os.Getenv("MAILER_FROM"),
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
		os.Getenv("MAILER_USERNAME"),
		os.Getenv("MAILER_PASSWORD"),
	)

	// SSL and TLS config
	ssl, err := strconv.ParseBool(os.Getenv("MAILER_SSL"))
	if err != nil {
		return errors.New("Cannot parse MAILER_SSL into bool")
	}

	startTLS, err := strconv.ParseBool(os.Getenv("MAILER_START_TLS"))
	if err != nil {
		return errors.New("Cannot parse MAILER_START_TLS into bool")
	}

	d.SSL = ssl
	d.TLSConfig = &tls.Config{
		InsecureSkipVerify: startTLS,
	}

	// Send email
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
