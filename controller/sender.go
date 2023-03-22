package controller

import "github.com/kwamekyeimonies/Golang-SMTP/model"


func NewGmailSender(name string,fromEmailAddress string,fromEmailPassword string) model.EmailSender{
	return &model.GmailSender{
		Name: name,
		FromEmailAddress: fromEmailAddress,
		FromEmailPassword: fromEmailPassword,
	}
}

