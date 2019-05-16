package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
)

func handle(req *events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       "Hello " + req.Body,
	}, nil

}

func main() {
	lambda.Start(handle)
}
