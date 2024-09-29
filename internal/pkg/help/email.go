package help

import (
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"os"
	"strings"
)

type Params struct {
	From     string
	Password string
	To       string
	Message  string
	Code     string
	UserName string
}

func SendVerificationCode(params Params) error {
	// Check working directory
	wd, _ := os.Getwd()
	log.Println("Current working directory:", wd)

	// Attempt to read HTML template
	htmlFile, err := os.ReadFile("internal/pkg/help/format.html")
	if err != nil {
		log.Println("Cannot read HTML file, using default template:", err)
		htmlFile = []byte(`<!DOCTYPE html>
		<html lang="en">
		<head>
		    <meta charset="UTF-8">
		    <title>Email Notification</title>
		</head>
		<body style="font-size: 20px;">
		    <p><strong>{{ .Message }}:</strong></p>
		    <p>Message: <strong>{{ .Code }}</strong></p>
		</body>
		</html>`)
	}

	// Parse HTML template
	temp, err := template.New("email").Parse(string(htmlFile))
	if err != nil {
		log.Println("Cannot parse HTML template:", err)
		return fmt.Errorf("failed to parse HTML template: %w", err)
	}

	// Execute the template with params
	var builder strings.Builder
	err = temp.Execute(&builder, params)
	if err != nil {
		log.Println("Cannot execute HTML template:", err)
		return fmt.Errorf("failed to execute HTML template: %w", err)
	}

	// Create the email message
	message := "From: " + params.From + "\n" +
		"To: " + params.To + "\n" +
		"Subject: Notification\n" +
		"MIME-Version: 1.0\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\n\n" +
		builder.String()

	// Send email using SMTP
	auth := smtp.PlainAuth("", params.From, params.Password, "smtp.gmail.com")
	err = smtp.SendMail("smtp.gmail.com:587", auth, params.From, []string{params.To}, []byte(message))
	if err != nil {
		log.Println("Could not send an email:", err)
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
