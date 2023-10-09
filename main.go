package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/google/uuid"
	"github.com/vineela-devarashetty/user-microservice/helper"
	"github.com/vineela-devarashetty/user-microservice/model"
	"github.com/vineela-devarashetty/user-microservice/service"
	"log"
	"net/http"
)

var dynamoDB *dynamodb.DynamoDB

func main() {
	// Initialize AWS session and DynamoDB client
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"), // Change to your AWS region
	})
	if err != nil {
		log.Fatal("Error creating AWS session:", err)
	}

	dynamoDB = dynamodb.New(sess)
	lambda.Start(handler)
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch request.HTTPMethod {
	case "POST":
		return createUser(ctx, request)
	case "GET":
		return getUser(ctx, request)
	case "PUT":
		return updateUser(ctx, request)
	case "DELETE":
		return deleteUser(ctx, request)
	default:
		return clientError(http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func createUser(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var user model.User
	err := json.Unmarshal([]byte(request.Body), &user)
	if err != nil {
		return clientError(http.StatusBadRequest, "Invalid request payload")
	}

	// Validate user data

	if err := helper.ValidateUser(user); err != nil {
		return clientError(http.StatusBadRequest, "Required fields are missing")
	}

	if !helper.IsValidEmail(user.Email) {
		return clientError(http.StatusBadRequest, "Invalid email format")
	}

	if !helper.IsValidDOB(user.DOB) {
		return clientError(http.StatusBadRequest, "Invalid date of birth format (YYYY-MM-DD)")
	}

	userID := uuid.New().String()

	user.UserID = userID

	err = service.CreateUser(ctx, &user, dynamoDB)
	if err != nil {
		return serverError(err)
	}

	return clientResponse(http.StatusCreated, user)
}

func getUser(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	userID := request.PathParameters["UserID"]

	fetchedUser, err := service.GetUser(ctx, userID, dynamoDB)
	if err != nil {
		if err.Error() == "User not found" {
			return clientError(http.StatusNotFound, "User not found")
		}
		return serverError(err)
	}

	return clientResponse(http.StatusOK, fetchedUser)
}

func updateUser(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	userID := request.PathParameters["UserID"]

	var updatedUser model.User
	err := json.Unmarshal([]byte(request.Body), &updatedUser)
	if err != nil {
		return clientError(http.StatusBadRequest, "Invalid request payload")
	}

	// Validate user data

	if err := helper.ValidateUser(updatedUser); err != nil {
		return clientError(http.StatusBadRequest, "Required fields are missing")
	}

	if !helper.IsValidEmail(updatedUser.Email) {
		return clientError(http.StatusBadRequest, "Invalid email format")
	}

	if !helper.IsValidDOB(updatedUser.DOB) {
		return clientError(http.StatusBadRequest, "Invalid date of birth format (YYYY-MM-DD)")
	}

	err = service.UpdateUser(ctx, userID, &updatedUser, dynamoDB)
	if err != nil {
		return serverError(err)
	}

	return clientResponse(http.StatusOK, updatedUser)
}

func deleteUser(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	userID := request.PathParameters["UserID"]

	err := service.DeleteUser(ctx, userID, dynamoDB)
	if err != nil {
		return serverError(err)
	}

	return clientResponse(http.StatusNoContent, nil)
}

func clientError(status int, message string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       fmt.Sprintf(`{"error": "%s"}`, message),
	}, nil
}

func serverError(err error) (events.APIGatewayProxyResponse, error) {
	log.Println(err.Error())
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       `{"error": "Internal Server Error"}`,
	}, nil
}

func clientResponse(status int, body interface{}) (events.APIGatewayProxyResponse, error) {
	responseBody, err := json.Marshal(body)
	if err != nil {
		return serverError(errors.New("Error marshalling response"))
	}
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       string(responseBody),
	}, nil
}
