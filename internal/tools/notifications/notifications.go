package notifications

import (
	"fmt"
	"time"

	"github.com/martinlindhe/notify"
)

var (
	cleaningHour, cleaningMinute int
	flagHour, flagMinute         int
)

func Start() {
	fmt.Println("Notification service started...")

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for now := range ticker.C {
		info, notify := isNotificationTime(now)

		if notify {
			sendNotification(info)
		}
	}
}

func isNotificationTime(now time.Time) (*Notification, bool) {
	allNotifications := generateNotifications(now)

    fmt.Println(allNotifications)

	for _, entry := range allNotifications {
		if now.Hour() == entry.Hour && now.Minute() == entry.Minute {
			return &entry.Notification, true
		}
	}

	return nil, false
}

func sendNotification(notification *Notification) {
	notify.Notify("Υπηρεσία ΓΕΠ", notification.Title, notification.Description, "")

	go playMP3(notification.Audio)
}
