package notification

import (
	"context"
	"errors"

	db "github.com/GO-notification-engine/db/sqlc"
)

type ServiceInterface interface {
}

type Service struct {
	repo RepositoryInterface
}

func NewService(repo RepositoryInterface) ServiceInterface {
	return &Service{repo: repo}
}

func (s *Service) CreateNotification(ctx context.Context, req CreateNotificationRequest) (NotificationResponse, error) {
	arg := db.CreateNotificationParams{
		Recipient:   req.Recipient,
		Message:     req.Message,
		ChannelType: req.ChannelType,
		Status:      req.Status,
	}

	notification, err := s.repo.CreateNotification(ctx, arg)
	if err != nil {
		return NotificationResponse{}, errors.New("error creating notification: " + err.Error())
	}

	result := NotificationResponse{
		ID:          notification.ID,
		Recipient:   notification.Recipient,
		Message:     notification.Message,
		ChannelType: notification.ChannelType,
		Status:      notification.Status,
		CreatedAt:   notification.CreatedAt,
		UpdatedAt:   notification.UpdatedAt,
	}

	return result, nil
}
