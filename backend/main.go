package main

import (
	"context"
	"encoding/base64"
	"net/http"
	"net/url"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/labstack/gommon/log"
)

type ContactFormData struct {
	FromEmail   string `json:"from_email"`
	MessageBody string `json:"message_body"`
}

func processFormData(body string) string {
	values, _ := url.ParseQuery(body)
	email := values["email"][0]
	name := values["name"][0]
	message := values["message"][0]
	next := values["_next"][0]
	log.Info("email = " + email)
	log.Info("name = " + name)
	log.Info("message = " + message)
	log.Info("next = " + next)

	sendMail(name, email, message)
	return next
}

// Handler handles API requests
func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {

	m := req.RequestContext.HTTP.Method
	p := req.RequestContext.HTTP.Path
	var b string

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

	log.Info("body = " + b)

	if p == "/contact" && m == "POST" {
		next := processFormData(b)
		//c := ContactFormData{}
		//json.Unmarshal([]byte(req.Body), &c)
		//log.Info("Email addr = " + c.FromEmail)

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
