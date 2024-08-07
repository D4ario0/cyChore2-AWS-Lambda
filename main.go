package main

import (
	"context"
	"cyChore2/emailsender"
	"cyChore2/utils"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	Users []emailsender.User `json:"users"`
}

func sendReminders(users *Event, sender *emailsender.EmailServer) {
	today := time.Now()
	weekOffset := utils.GetWeeksUntil(today)

	for _, user := range users.Users {
		user.AssignTasks(weekOffset)
		err := sender.SendWeeklyTaskMail(&user, "static/notification.html")
		if err != nil {
			log.Printf("Error sending email to %s: %v", user.Email, err)
		} else {
			log.Printf("Task reminder sent to %s at %v", user.Email, today)
		}
	}
}

func HandleRequest(ctx context.Context, users Event) (int, error) {
	if len(users.Users) == 0 {
		return http.StatusBadRequest, fmt.Errorf("received empty users list")
	}

	sender := emailsender.EmailServer{
		Domain: os.Getenv("SMTP_DOMAIN"),
		Port:   os.Getenv("SMTP_PORT"),
		Email:  os.Getenv("SMTP_EMAIL"),
		Passwd: os.Getenv("SMTP_PSSWD"),
	}

	sendReminders(&users, &sender)
	return http.StatusAccepted, nil
}

func main() {
	lambda.Start(HandleRequest)
}
