package notification

import "time"

type CreateNotificationRequest struct {
	Recipient   string `json:"recipient"`
	Message     string `json:"message"`
	ChannelType string `json:"channel_type"`
	Status      string `json:"status"`
}

type NotificationResponse struct {
	ID          int64     `json:"id"`
	Recipient   string    `json:"recipient"`
	Message     string    `json:"message"`
	ChannelType string    `json:"channel_type"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
