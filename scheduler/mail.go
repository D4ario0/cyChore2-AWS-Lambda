package scheduler

import (
	"bytes"
	"fmt"
	"io"
	"lambda-cychore/types"
	"net/http"
	"os"
	"text/template"

	"gopkg.in/gomail.v2"
)

const (
	subject              = "Your Weekly Task Reminder"
	errFetchTemplate     = "failed to fetch email template: %w"
	errParseTemplate     = "failed to parse email template: %w"
	errExecuteTemplate   = "failed to execute template: %w"
	errFetchHTMLTemplate = "failed to fetch HTML template: %w"
	errReadHTMLTemplate  = "failed to read HTML template: %w"
)

func CreateEmailSender() (types.UserProcessor, error) {
	htmlTemplate, err := fetchHTMLTemplate(os.Getenv("EMAIL_TEMPLATE"))
	if err != nil {
		return nil, fmt.Errorf(errFetchTemplate, err)
	}

	tmpl, err := template.New("emailTemplate").Parse(htmlTemplate)
	if err != nil {
		return nil, fmt.Errorf(errParseTemplate, err)
	}

	return func(user *types.User) error {
		userInfo := struct {
			Name string
			Task string
		}{
			Name: user.Name,
			Task: user.Task,
		}

		parsedHTML, err := parseHTML(tmpl, userInfo)
		if err != nil {
			return err
		}

		err = sendEmail(parsedHTML, user.Email)
		return err
	}, nil
}

func fetchHTMLTemplate(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf(errFetchHTMLTemplate, err)
	}
	defer resp.Body.Close()

	htmlBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf(errReadHTMLTemplate, err)
	}

	return string(htmlBytes), nil
}

func parseHTML(tmpl *template.Template, data interface{}) (string, error) {
	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		return "", fmt.Errorf(errExecuteTemplate, err)
	}
	return body.String(), nil
}

func sendEmail(emailBody string, receiver string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("SMTP_EMAIL"))
	m.SetHeader("To", string(receiver))
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", emailBody)

	d := gomail.NewDialer(
		os.Getenv("SMTP_DOMAIN"),
		587,
		os.Getenv("SMTP_EMAIL"),
		os.Getenv("SMTP_PASSWORD"),
	)

	return d.DialAndSend(m)
}
