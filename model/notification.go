package model

// Notification is our message structure that will help to describe the Notification"
type Notification struct {
	ID         string             `json:"id,omitempty"`
	Message    string             `json:"message"`
	Stacktrace string             `json:"stacktrace,omitempty"`
	Level      *NotificationLevel `json:"level"`
}

// NotificationLevel is the structure that will help describe the level of Notification
type NotificationLevel struct {
	Name     string `json:"name"`
	Bloquant bool   `json:"bloquant"`
}

var Notifications []Notification
