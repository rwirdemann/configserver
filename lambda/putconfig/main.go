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
	if err != nil {
		log.Printf("error: %v", err)
		return buildResponse(http.StatusBadRequest)
	}

	err = dynamo.AddConfigItem(item)
	if err != nil {
		log.Printf("error: %v", err)
		return buildResponse(http.StatusInternalServerError)
	}

	log.Printf("created or updated config item %v", item)
	return buildResponse(http.StatusNoContent)
}

func buildResponse(httpStatus int) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: httpStatus,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Credentials": "true",
		},
	}

}
