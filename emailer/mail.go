package emailer

import "fmt"

const (
	emailSubject  = "Your weekly task reminder"
	emailTemplate = `<html lang="en">
<body style="font-family: system-ui, -apple-system;">
	<h4>Hi <strong>%s</strong>,</h4>
	<p>Your assigned task for this week is:</p>
	<p<em>%s</em></p>
	<p>If you have any questions, do not reach out.</p>
	<p>Cold regards,</p>
	<p><em><strong>The Task Management Team</strong></em></p>
</body>
</html>`
)

func ParseEmailContent(to, name, task string) []byte {
	body := fmt.Sprintf(emailTemplate, name, task)

	msg := fmt.Sprintf("To: %s\r\n"+
		"Subject: %s\r\n"+
		"MIME-Version: 1.0\r\n"+
		"Content-Type: text/html; charset=\"UTF-8\"\r\n"+
		"\r\n"+
		"%s", to, emailSubject, body)

	return []byte(msg)
}
