# user-microservice
Golang implementation of User microservices

# User Microservice

This is a CRUD (Create, Read, Update, Delete) microservice for managing user records, deployed on AWS Lambda and accessible via AWS API Gateway. The service is developed in GoLang and utilizes AWS DynamoDB as the data repository.

## Table of Contents

- [API Endpoints](#api-endpoints)
- [Error Handling](#error-handling)
- [Deployment](#deployment)
- [Testing](#testing)
- [Getting Started](#getting-started)

## API Endpoints

### Create User
- **Endpoint:** `/users`
- **HTTP Method:** POST
- **Description:** Adds a new user with fields UserID, Name, Email and DOB.

### Read User
- **Endpoint:** `/users/{UserID}`
- **HTTP Method:** GET
- **Description:** Retrieves user information based on UserID.

### Update User
- **Endpoint:** `/users/{UserID}`
- **HTTP Method:** PUT
- **Description:** Modifies existing user details using UserID.

### Delete User
- **Endpoint:** `/users/{UserID}`
- **HTTP Method:** DELETE
- **Description:** Removes a user record based on UserID.

## Error Handling

The microservice implements thorough error handling. It returns descriptive error messages and suitable HTTP status codes for various types of errors.

## Deployment

To deploy this microservice, we use the AWS Serverless framework. Follow these steps:

1. Clone this repository:
   ```shell
   git clone <repository_url>
2. Build the project using the script ./start.sh
3. Install Serverless framework (if not already installed):
    npm install -g serverless
4. Configure your AWS credentials:
    serverless config credentials --provider aws --key <your_aws_access_key> --secret <your_aws_secret_key>
5. Deploy the service to AWS Lambda:
    serverless deploy
6. Note the deployed API Gateway endpoint URL provided in the output.

# Testing
This microservice includes unit tests to ensure its functionality. It uses a mock framework to simulate DynamoDB during tests. To run the tests, follow these steps:

1. Install GoLang if not already installed: [GoLang Installation](https://go.dev/doc/install)

2. Install required Go packages:
    go mod tidy
3. Run the unit tests:
    go test ./...

# Integration Testing
1. Access the AWS API Gateway Console:

    1. Open your web browser and go to the [AWS Management Console](https://us-east-1.console.aws.amazon.com).
    
    2. Sign in to your AWS account if you're not already logged in.
    
    3. In the AWS Management Console, navigate to the "Services" menu and   select "API Gateway" under the "Networking & Content Delivery" section.
    Select Your API:

2. In the AWS API Gateway Console, select the API that corresponds to your deployed User Microservice.

3. Testing the Endpoints:

    1. Click on the endpoint you want to test.

    2. After clicking on the endpoint, you'll see a page where you can test the endpoint using the API Gateway Console.
    3. Click the "Test" button or "Invoke" button (the exact label may vary depending on your API setup).
    4. For POST and PUT request, provide the required data in the request body.
    5. Click the "Test" button to send the request to your User Microservice.
    6. You will receive the response from your User Microservice, including the HTTP status code and response body, if applicable.

Below are the example requests for the User Microservice POST endpoint

{
    "Name": "John Doe",
    "Email" : "john.doe@gmail.com",
    "DOB": "1982-12-12"
}