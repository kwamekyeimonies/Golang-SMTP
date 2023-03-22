package main

import (
	"log"

	"github.com/kwamekyeimonies/Golang-SMTP/controller"
	"github.com/kwamekyeimonies/Golang-SMTP/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Printf("Error: %v", err.Error())
	}
	sender := controller.NewGmailSender(config.EmailSenderName, config.EmailSenderEmail, config.EmailPassword)
	subject := "Testing Email SMTP"
	content := `
	<h1>Hello,Dear</h1>
	<p>This is a test message from Quantum_Monies@001 </p>
	`
	to := []string{"tenkorangd5@gmail.com"}
	attachfiles := []string{"./ReadMe.md"}

	err = sender.SendEmail(subject, content, to, nil, nil, attachfiles)
	if err != nil {
		log.Println(err.Error())
	}

}
