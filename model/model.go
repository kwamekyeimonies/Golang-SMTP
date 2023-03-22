package model

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

const (
	SmtpAuthAddress   = "smtp.gmail.com"
	SmtpServerAddress = "smtp.gmail.com:587"
)

type EmailSender interface {
	SendEmail(
		Subject string,
		Content string,
		To []string,
		Cc []string,
		Bcc []string,
		AttachFiles []string,
	) error
}

type GmailSender struct {
	Name              string
	FromEmailAddress  string
	FromEmailPassword string
}

func (sender *GmailSender) SendEmail(
	Subject string,
	Content string,
	To []string,
	Cc []string,
	Bcc []string,
	AttachFiles []string,
) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", sender.Name, sender.FromEmailAddress)
	e.Subject = Subject
	e.HTML = []byte(Content)
	e.To = To
	e.Cc = Cc
	e.Bcc = Bcc

	for _, f := range AttachFiles {
		_, err := e.AttachFile(f)
		if err != nil {
			return fmt.Errorf("Failed to Attach file %s: %w", f, err)
		}
	}

	smtpAuth := smtp.PlainAuth("", sender.FromEmailAddress, sender.FromEmailPassword, SmtpAuthAddress)
	return e.Send(SmtpServerAddress, smtpAuth)
}
