package notifications

import "log"

type PushNotification struct {
	Phone string
}

func NewPushTypeNotification(phone string) INotifications {
	return PushNotification{
		Phone: phone,
	}
}

func (pn PushNotification) NotifyUser() error {
	log.Println("User successfully notitfied via phone: " + pn.Phone)

	return nil
}