package service

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/vineela-devarashetty/user-microservice/model"
)

// MockDynamoDB is a mock implementation of the DynamoDBAPI interface.
type MockDynamoDB struct {
	mock.Mock
}

// Create a mock DynamoDB instance that implements the DynamoDBAPI interface.
func NewMockDynamoDB() *MockDynamoDB {
	return &MockDynamoDB{}
}

func (m *MockDynamoDB) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	args := m.Called(input)
	return args.Get(0).(*dynamodb.PutItemOutput), args.Error(1)
}

func (m *MockDynamoDB) GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	args := m.Called(input)
	return args.Get(0).(*dynamodb.GetItemOutput), args.Error(1)
}

func (m *MockDynamoDB) UpdateItem(input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
	args := m.Called(input)
	return args.Get(0).(*dynamodb.UpdateItemOutput), args.Error(1)
}

func (m *MockDynamoDB) DeleteItem(input *dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error) {
	args := m.Called(input)
	return args.Get(0).(*dynamodb.DeleteItemOutput), args.Error(1)
}

func TestCreateUser(t *testing.T) {
	// Create a mock DynamoDB client
	mockDynamoDB := NewMockDynamoDB()

	user := &model.User{
		UserID: "123",
		Name:   "John Doe",
		Email:  "john.doe@example.com",
		DOB:    "1990-01-01",
	}

	// Mock the PutItem operation to return success
	mockDynamoDB.On("PutItem", mock.Anything).Return(&dynamodb.PutItemOutput{}, nil)

	err := CreateUser(context.Background(), user, mockDynamoDB)
	assert.NoError(t, err)

	// Verify that the PutItem method was called with the expected input
	mockDynamoDB.AssertCalled(t, "PutItem", mock.AnythingOfType("*dynamodb.PutItemInput"))
}

func TestGetUser(t *testing.T) {
	// Create a mock DynamoDB client
	mockDynamoDB := NewMockDynamoDB()

	userID := "123"
	user := &model.User{
		UserID: "123",
		Name:   "John Doe",
		Email:  "john.doe@example.com",
		DOB:    "1990-01-01",
	}

	// Mock the GetItem operation to return the user
	mockDynamoDB.On("GetItem", mock.Anything).Return(&dynamodb.GetItemOutput{
		Item: map[string]*dynamodb.AttributeValue{
			"UserID": {S: &user.UserID},
			"Name":   {S: &user.Name},
			"Email":  {S: &user.Email},
			"DOB":    {S: &user.DOB},
		},
	}, nil)

	result, err := GetUser(context.Background(), userID, mockDynamoDB)
	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}

	if result == nil {
		t.Fatal("Expected non-nil user, but got nil")
	}
	assert.NoError(t, err)
	assert.Equal(t, user, result)

	// Verify that the GetItem method was called with the expected input
	mockDynamoDB.AssertCalled(t, "GetItem", mock.AnythingOfType("*dynamodb.GetItemInput"))
}

func TestUpdateUser(t *testing.T) {
	// Create a mock DynamoDB client
	mockDynamoDB := NewMockDynamoDB()

	userID := "123"
	updatedUser := &model.User{
		UserID: "123",
		Name:   "Updated Name",
		Email:  "updated.email@example.com",
		DOB:    "1991-02-02",
	}

	// Mock the UpdateItem operation to return success
	mockDynamoDB.On("UpdateItem", mock.Anything).Return(&dynamodb.UpdateItemOutput{}, nil)

	err := UpdateUser(context.Background(), userID, updatedUser, mockDynamoDB)
	assert.NoError(t, err)

	// Verify that the UpdateItem method was called with the expected input
	mockDynamoDB.AssertCalled(t, "UpdateItem", mock.AnythingOfType("*dynamodb.UpdateItemInput"))
}

func TestDeleteUser(t *testing.T) {
	// Create a mock DynamoDB client
	mockDynamoDB := NewMockDynamoDB()

	userID := "123"

	// Mock the DeleteItem operation to return success
	mockDynamoDB.On("DeleteItem", mock.Anything).Return(&dynamodb.DeleteItemOutput{}, nil)

	err := DeleteUser(context.Background(), userID, mockDynamoDB)
	assert.NoError(t, err)

	// Verify that the DeleteItem method was called with the expected input
	mockDynamoDB.AssertCalled(t, "DeleteItem", mock.AnythingOfType("*dynamodb.DeleteItemInput"))
}
