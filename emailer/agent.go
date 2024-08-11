package emailer

import (
	"fmt"
	"net/smtp"
)

type EmailAgent struct {
	Domain string
	Port   string
	Email  string
	Passwd string
}

func (sender EmailAgent) SendWeeklyTaskMail(u *User) error {
	auth := smtp.PlainAuth("", sender.Email, sender.Passwd, sender.Domain)
	host := fmt.Sprintf("%s:%s", sender.Domain, sender.Port)

	msg := ParseEmailContent(u.Email, u.Name, u.Task)

	err := smtp.SendMail(host, auth, sender.Email, []string{u.Email}, msg)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
