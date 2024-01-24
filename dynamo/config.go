package dynamo

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"log"
	"os"
)

const tableName = "ConfigItems"

type ConfigItem struct {
	Key   string
	Value string
}

func AddConfigItem(item ConfigItem) error {
	s, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("REGION"))},
	)
	if err != nil {
		return err
	}
	db := dynamodb.New(s)
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}
	_, err = db.PutItem(input)

	if err != nil {
		return err
	}
	log.Printf("config item added or updated %v", item)
	return nil
}

func DeleteConfigItem(item ConfigItem) error {
	s, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("REGION"))},
	)
	if err != nil {
		return err
	}
	db := dynamodb.New(s)
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"Key": {
				N: aws.String(item.Key),
			},
		},
		TableName: aws.String(tableName),
	}

	_, err = db.DeleteItem(input)
	if err != nil {
		log.Fatalf("Got error calling DeleteItem: %s", err)
	}
	return nil
}
