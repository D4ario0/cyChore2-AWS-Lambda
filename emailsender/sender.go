package emailsender

import (
	"bytes"
	"fmt"
	"net/smtp"
	"text/template"
)

type EmailServer struct {
	Domain string
	Port   string
	Email  string
	Passwd string
}

func (sender EmailServer) SendWeeklyTaskMail(u *User, templatePath string) error {
	auth := smtp.PlainAuth("", sender.Email, sender.Passwd, sender.Domain)
	host := fmt.Sprintf("%s:%s", sender.Domain, sender.Port)
	subject := "Your weekly task reminder"

	body, err := getEmailBody(templatePath, u.Username, u.Task)
	if err != nil {
		return fmt.Errorf("failed to get email body: %w", err)
	}

	// Construct the email message
	msg := []byte(fmt.Sprintf("To: %s\r\n"+
		"Subject: %s\r\n"+
		"MIME-Version: 1.0\r\n"+
		"Content-Type: text/html; charset=\"UTF-8\"\r\n"+
		"\r\n"+
		"%s", u.Email, subject, body))

	err = smtp.SendMail(host, auth, sender.Email, []string{u.Email}, msg)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}

func getEmailBody(templatePath, username, task string) ([]byte, error) {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return nil, fmt.Errorf("error parsing template: %w", err)
	}

	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, struct {
		Name       string
		AssignTask string
	}{
		Name:       username,
		AssignTask: task,
	})
	if err != nil {
		return nil, fmt.Errorf("error executing template: %w", err)
	}

	return buffer.Bytes(), nil
}
