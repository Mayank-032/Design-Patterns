package main

import (
	"errors"
	"factory-pattern/notifications"
	"log"
)

const (
	EventTypeUserRegistration = "user-registration"
    EventTypeOrderShipped     = "order-shipped"
)

func fetchNotificationMedium(eventType string) (notifications.INotifications, error)  {
	switch eventType {
	case "user-registration":
		notification := notifications.NewEmailTypeNotification("m*************m", "user registered successfully")
		return notification, nil
	case "order-shipped":
		notification := notifications.NewPushTypeNotification("97XXXXXXXX")
		return notification, nil
	default:
		log.Println("Invalid event type")
		return nil, errors.New("invalid event type")
	}
}

func main() {
	notificationMedium, err := fetchNotificationMedium(EventTypeOrderShipped)
	if err != nil {
		log.Fatalln("exiting on failure...")
	}

	err = notificationMedium.NotifyUser()
	if err != nil {
		log.Println("unable to notify user")
		log.Fatalln("exiting on failure...")
		return
	}

	log.Println("successfully notified user")
	log.Fatalln("exiting on success...")
}
