package main

import (
	"context"
	"encoding/base64"
	"net/http"
	"net/url"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
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
	}

	// Attempt to send the email.
	_, err = svc.SendEmail(input)

	if err != nil {
		log.Info("Error sending email: " + err.Error())
	}
}

func processFormData(body string) string {
	values, _ := url.ParseQuery(body)
	email := values["email"][0]
	name := values["name"][0]
	message := values["message"][0]
	next := values["_next"][0]

	sendMail(name, email, message)
	return next
}

func extractRequestData(req events.APIGatewayV2HTTPRequest) (m, p, b string) {
	m = req.RequestContext.HTTP.Method
	p = req.RequestContext.HTTP.Path

	if req.IsBase64Encoded {
		decoded, err := base64.StdEncoding.DecodeString(req.Body)
		if err != nil {
			log.Info("Error decoding base64 encoded string.")
			b = "nil"
		} else {
			b = string(decoded)
		}
	} else {
		b = req.Body
	}

	return
}

// Handler handles API requests
func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {

	m, p, b := extractRequestData(req)
	log.Info("body = " + b)

	if p == "/contact" && m == "POST" {
		// the return value is where we redirect to...
		next := processFormData(b)

		h := make(map[string]string)
		h["Location"] = next

		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusFound,
			Headers:    h,
			Body:       `{"message": "Will respond to email addr asap..."}`,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusNotFound,
		Body:       `{"message": "Resource not found on this endpoint."}`,
	}, nil

}

func main() {
	lambda.Start(Handler)
}
