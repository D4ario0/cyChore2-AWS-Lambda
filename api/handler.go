package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"lambda-cychore/scheduler"
	"lambda-cychore/types"

	"github.com/aws/aws-lambda-go/events"
)

const (
	errInvalidRequestBody    = "Invalid request body"
	errFailedToCreateSender  = "Failed to create email sender: %v"
	errFailedToMarshalResult = "Failed to marshal response: %v"
)

// Response structure to be used in the API response.
type APIResponse struct {
	Message string   `json:"message"`
	Errors  []string `json:"errors,omitempty"`
}

// Handler is the main Lambda function handler.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var users types.UserList

	// Unmarshal the request body into the UserList struct
	err := json.Unmarshal([]byte(request.Body), &users)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       fmt.Sprintf(`{"error": "%s"}`, errInvalidRequestBody),
		}, err
	}

	// Assign tasks to users
	users.ForEach(scheduler.AssignTasks)

	// Create the email sender
	emailSender, err := scheduler.CreateEmailSender()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       fmt.Sprintf(errFailedToCreateSender, err),
		}, nil
	}

	// Use the ForEach method with the email sender
	errors := users.ForEach(emailSender)

	// Prepare the response message
	responseMsg := fmt.Sprintf("Processed %d users", len(users.Users))
	if len(errors) > 0 {
		responseMsg += fmt.Sprintf(", encountered %d errors", len(errors))
	}

	// Build the response structure
	response := APIResponse{
		Message: responseMsg,
		Errors:  make([]string, len(errors)),
	}

	// Populate the errors in the response
	for i, err := range errors {
		response.Errors[i] = err.Error()
	}

	// Marshal the response to JSON
	jbytes, err := json.Marshal(response)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       fmt.Sprintf(`{"error": "%s"}`, fmt.Sprintf(errFailedToMarshalResult, err)),
		}, err
	}

	// Return the successful response
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(jbytes),
	}, nil
}
