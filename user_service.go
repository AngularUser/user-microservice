package service

import (
	"context"
	"errors"
	// "fmt"
	// "encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	// "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	// "github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/vineela-devarashetty/user-microservice/model"
	// "log"
)

var tableName = "Users" // DynamoDB table name

// Create a new user
func CreateUser(ctx context.Context, user *model.User, dynamoDB DynamoDBAPI) error {

	// Create a new user record in DynamoDB
	putItemInput := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: map[string]*dynamodb.AttributeValue{
			"UserID": {S: aws.String(user.UserID)},
			"Name":   {S: aws.String(user.Name)},
			"Email":  {S: aws.String(user.Email)},
			"DOB":    {S: aws.String(user.DOB)},
			// Add other fields as needed
		},
	}

	_, err := dynamoDB.PutItem(putItemInput)
	if err != nil {
		return err
	}

	return nil
}

// Read user details by UserID
func GetUser(ctx context.Context, userID string, dynamoDB DynamoDBAPI) (*model.User, error) {
	// Retrieve user record from DynamoDB
	getItemInput := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"UserID": {S: aws.String(userID)},
		},
	}

	result, err := dynamoDB.GetItem(getItemInput)
	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, errors.New("User not found")
	}
	// log.Println("result item is :", result.Item)

	// var userString string
	var user model.User
	// attributevalue.Unmarshal(result.Item, &userString)

	// json.Unmarshal([]byte(userString), &user)

	err = unmarshalDynamoDBItem(result.Item, &user)
	// log.Println("user is :", user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Update user details by UserID
func UpdateUser(ctx context.Context, userID string, updatedUser *model.User, dynamoDB DynamoDBAPI) error {
	// Validate user data here if needed

	// Update user record in DynamoDB
	updateItemInput := &dynamodb.UpdateItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"UserID": {S: aws.String(userID)},
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":Name":  {S: aws.String(updatedUser.Name)},
			":Email": {S: aws.String(updatedUser.Email)},
			":DOB":   {S: aws.String(updatedUser.DOB)},
			// Add other fields as needed
		},
		UpdateExpression: aws.String("SET Name = :name, Email = :email, DOB = :dob"),
	}

	_, err := dynamoDB.UpdateItem(updateItemInput)
	if err != nil {
		return err
	}

	return nil
}

// Delete user by UserID
func DeleteUser(ctx context.Context, userID string, dynamoDB DynamoDBAPI) error {
	// Delete user record from DynamoDB
	deleteItemInput := &dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"UserID": {S: aws.String(userID)},
		},
	}

	_, err := dynamoDB.DeleteItem(deleteItemInput)
	if err != nil {
		return err
	}

	return nil
}

func unmarshalDynamoDBItem(attribute map[string]*dynamodb.AttributeValue, out interface{}) error {
	// for key, value := range attribute {
	//     log.Println("Key:", key) // Use fmt.Printf or log.Println to print the content
	// 	log.Println("Value:", value)
	// }
	// av, err := dynamodbattribute.MarshalMap(attribute)
	// if err != nil {
	// 	log.Println("err s :", err)
	// 	return err
	// }
	// log.Println("av is :", av)
	return dynamodbattribute.UnmarshalMap(attribute, out)
}
