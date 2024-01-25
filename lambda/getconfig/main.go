package main

import (
	"errors"
	"github.com/rwirdemann/configserver/dynamo"
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
	item, err := dynamo.GetConfigItem(key)
	if err == nil {
		log.Printf("found config item for key '%s': %s", key, item.Value)
		return events.APIGatewayProxyResponse{
			Body:       item.Value,
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin":      "*",
				"Access-Control-Allow-Credentials": "true",
			},
		}
	}
	status := http.StatusInternalServerError
	if errors.Is(err, dynamo.NotFound) {
		log.Printf("config item for key '%s' not found", key)
		status = http.StatusNotFound
	} else {
		log.Printf("error fetching config item for key: '%s': %v", key, err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Credentials": "true",
		},
	}
}
