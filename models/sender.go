package models

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

type EmailSender interface {
	SendEmail(
		subject string,
		content string,
		to []string,
		cc []string,
		bcc []string,
		attachFile []string,
	) error
}

type GmailSender struct {
	Name              string
	FromEmailAddress  string
	FromEmailPassword string
}

func NewGmailSender(Name string, FromEmailAddress string, FromEmailPassword string) EmailSender {
	return &GmailSender{
		Name:              Name,
		FromEmailAddress:  FromEmailAddress,
		FromEmailPassword: FromEmailPassword,
	}
}

func (sender *GmailSender) SendEmail(
	subject string,
	content string,
	to []string,
	cc []string,
	bcc []string,
	attachFile []string,
) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", sender.Name, sender.FromEmailAddress)
	e.To = to
	e.Cc = cc
	e.Bcc = bcc
	e.Subject = subject
	e.Text = []byte(content)
	e.HTML = []byte(content)
	for _, f := range attachFile {
		_, err := e.AttachFile(f)
		if err != nil {
			return fmt.Errorf("error when attaching file: %s", err.Error())
		}
	}

	auth := smtp.PlainAuth("", sender.FromEmailAddress, sender.FromEmailPassword, "smtp.gmail.com")
	return e.Send("smtp.gmail.com:587", auth)
}
