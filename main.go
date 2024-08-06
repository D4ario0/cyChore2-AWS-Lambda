package main

import (
	"context"
	"cyChore2/emailsender"
	"cyChore2/utils"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

type LambdaEvent map[string]interface{}

type LambdaResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func sendReminders(ctx context.Context, event LambdaEvent) (LambdaResponse, error) {
	statusCode := 200
	today := time.Now()
	weekOffset := utils.GetWeeksUntil(today)

	userList := emailsender.NewUserList()
	err := userList.ReadUsers("users.json")
	if err != nil {
		log.Printf("Error reading users: %v", err)
		return LambdaResponse{StatusCode: 500}, fmt.Errorf("failed to read users: %w", err)
	}

	sender := emailsender.EmailServer{
		Domain: os.Getenv("SMTP_DOMAIN"),
		Port:   os.Getenv("SMTP_PORT"),
		Email:  os.Getenv("SMTP_EMAIL"),
		Passwd: os.Getenv("SMTP_PASSWORD"),
	}

	failedEmails := 0
	for _, user := range userList.Users {
		user.AssignTasks(weekOffset)
		err := sender.SendWeeklyTaskMail(&user, "static/notification.html")
		if err != nil {
			log.Printf("Error sending email to %s: %v", user.Email, err)
			failedEmails++
		} else {
			log.Printf("Task reminder sent to %s at %v", user.Email, today)
		}
	}

	if failedEmails > 0 {
		statusCode = 503
		return LambdaResponse{
			StatusCode: statusCode,
			Message:    fmt.Sprintf("Completed with %d failed emails", failedEmails),
		}, nil
	}

	return LambdaResponse{
		StatusCode: statusCode,
		Message:    "All reminders sent successfully",
	}, nil
}

func main() {
	lambda.Start(sendReminders)
}
