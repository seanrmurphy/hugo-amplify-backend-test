package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/labstack/gommon/log"
)

type ContactFormData struct {
	FromEmail   string `json:"from_email"`
	MessageBody string `json:"message_body"`
}

//var e *echo.Echo
//var echoAdapter *echoadapter.EchoLambda
//var nullHandler = false

//// Lambda mode determines whether this is run locally or remotely
//var lambdaMode = true

// Handler handles API requests
func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {

	//r := core.RequestAccessor{}
	//httpRequest, err := r.ProxyEventToHTTPRequest(req)
	//eventString := fmt.Sprintf("Event = %v", req)
	//log.Info(eventString)
	//httpRequestString := fmt.Sprintf("Request = %v", httpRequest)
	log.Info("method = " + req.RequestContext.HTTP.Method)

	//if err != nil {
	//return events.APIGatewayProxyResponse{
	//StatusCode: http.StatusForbidden,
	//Body:       `{"message": "Error converting from event to http request."}`,
	//}, nil
	//}

	if req.RequestContext.HTTP.Method == "GET" {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusForbidden,
			Body:       `{"message": "HTTP GET method not supported on this endpoint."}`,
		}, nil
	}

	if req.RequestContext.HTTP.Method == "POST" {
		c := ContactFormData{}
		json.Unmarshal([]byte(req.Body), &c)
		log.Info("Email addr = " + c.FromEmail)

		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Body:       `{"message": "Will respond to email addr asap..."}`,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusForbidden,
		Body:       `{"message": "Unexpected HTTP method not supported on this endpoint."}`,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
