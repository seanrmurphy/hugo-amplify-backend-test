package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/labstack/gommon/log"
)

const (
	awsRegion = "eu-west-1"

	// should probably be somewhere in an env variable or otherwise not hardcoded...
	contactRecipientEmail = "sean@gopaddy.ch"

	sender = "sean@gopaddy.ch"

	CharSet = "UTF-8"
)

func sendMail(name, email, message string) {

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(awsRegion)},
	)

	svc := ses.New(sess)

	subject := "Contact Form Submitted by <" + email + "> " + name

	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{
				aws.String(email)},
			ToAddresses: []*string{
				aws.String(contactRecipientEmail),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(message),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(CharSet),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String(sender),
		// Uncomment to use a configuration set
		//ConfigurationSetName: aws.String(ConfigurationSet),
	}
	// Attempt to send the email.
	_, err = svc.SendEmail(input)

	if err != nil {
		log.Info("Error sending email: " + err.Error())
	}
}
