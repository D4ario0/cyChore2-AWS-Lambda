package emailsender

import (
	"fmt"
	"net/smtp"
)

type EmailServer struct {
	Domain string
	Port   string
	Email  string
	Passwd string
}

func (sender EmailServer) SendWeeklyTaskMail(u *User) error {
	auth := smtp.PlainAuth("", sender.Email, sender.Passwd, sender.Domain)
	host := fmt.Sprintf("%s:%s", sender.Domain, sender.Port)
	msg := []byte("")

	err := smtp.SendMail(host, auth, sender.Email, []string{u.Email}, msg)
	return err
}
