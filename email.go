package main
import(

	"net/smtp"
	"os"
	"fmt"
)

func sendmail(subject, body string) {
	from := "java99130@gmail.com"
	to := "trishanth.yuvaraj@gmail.com"
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	password := os.Getenv("SMTP_PASSWORD")

	auth := smtp.PlainAuth("", from, password, smtpHost)

	msg := fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\n\n%s", from, to, subject, body)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(msg))
	if err != nil {
		fmt.Println("Error sending email:", err)
	}
}