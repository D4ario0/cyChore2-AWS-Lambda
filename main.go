package main

import (
	"encoding/json"
	"fmt"
	"lambda-cychore/emailer"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	Message string `json:"message"`
}

func sendReminders(user *emailer.User, agent *emailer.EmailAgent) {
	today := time.Now()
	weekOffset := emailer.GetWeeksUntil(today)
	user.AssignTasks(weekOffset)

	err := agent.SendWeeklyTaskMail(user)
	if err != nil {
		log.Printf("Error sending email to %s: %v", user.Email, err)
	} else {
		log.Printf("Task reminder sent to %s at %v", user.Email, today)
	}

}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var user emailer.User

	err := json.Unmarshal([]byte(request.Body), &user)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest}, err
	}

	sender := emailer.EmailAgent{
		Domain: os.Getenv("SMTP_DOMAIN"),
		Port:   os.Getenv("SMTP_PORT"),
		Email:  os.Getenv("SMTP_EMAIL"),
		Passwd: os.Getenv("SMTP_PASSWORD"),
	}

	sendReminders(&user, &sender)

	mssg := Response{Message: fmt.Sprintf("%s and %s and %d should send email", user.Name, user.Email, user.Bufferindex)}

	jbytes, err := json.Marshal(mssg)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(jbytes),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
