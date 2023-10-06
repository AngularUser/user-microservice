package service

import "github.com/aws/aws-sdk-go/service/dynamodb"

// DynamoDBAPI defines a subset of DynamoDB operations.
type DynamoDBAPI interface {
	PutItem(*dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
	GetItem(*dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error)
	UpdateItem(*dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error)
	DeleteItem(*dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error)
}
