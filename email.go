package yonoma

import (
	"net/http"
)

// EmailRequest represents the structure for sending an email.
type EmailRequest struct {
	FromEmail    string `json:"from_email"`
	ToEmail      string `json:"to_email"`
	Subject      string `json:"subject"`
	MailTemplate string `json:"mail_template"`
}

// Send sends an email using the Yonoma API.
func (e *Client) Send(request EmailRequest) ([]byte, error) {
	endpoint := "/email/send"
	return e.request(http.MethodPost, endpoint, request)
}
