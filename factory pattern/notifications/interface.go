package notifications

type INotifications interface {
	NotifyUser() error
}