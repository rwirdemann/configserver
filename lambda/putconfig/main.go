package main

import (
	"encoding/json"
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
	item := dynamo.ConfigItem{}
	err := json.Unmarshal([]byte(r.Body), &item)
	log.Printf("config item %v", item)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Headers: map[string]string{
				"Access-Control-Allow-Origin":      "*",
				"Access-Control-Allow-Credentials": "true",
			},
		}
	}

	err = dynamo.AddConfigItem(item)
	if err != nil {
		log.Printf("error: %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers: map[string]string{
				"Access-Control-Allow-Origin":      "*",
				"Access-Control-Allow-Credentials": "true",
			},
		}
	}

	return events.APIGatewayProxyResponse{
		Body:       "https://77srys74sh.execute-api.eu-central-1.amazonaws.com/dev/jobs",
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Credentials": "true",
		},
	}
}
