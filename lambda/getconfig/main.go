package main

import (
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(func(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		return handler(r), nil
	})
}

func handler(r events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	key := r.PathParameters["key"]
	log.Printf("get request for key '%s'", key)

	return events.APIGatewayProxyResponse{
		Body:       "https://77srys74sh.execute-api.eu-central-1.amazonaws.com/dev/jobs",
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Credentials": "true",
		},
	}
}
