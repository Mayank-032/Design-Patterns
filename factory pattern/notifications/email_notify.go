package notifications

import "log"

type EmailNotification struct {
	Email   string
	Message string
}

func NewEmailTypeNotification(email, msg string) INotifications {
	return EmailNotification{
		Email: email,
		Message: msg,
	}
}

func (en EmailNotification) NotifyUser() error {
	log.Println("User successfully notitfied via email: " + en.Email)

	return nil
}