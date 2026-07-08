package notification

import (
	"context"

	db "github.com/GO-notification-engine/db/sqlc"
)

type RepositoryInterface interface {
	CreateNotification(ctx context.Context, arg db.CreateNotificationParams) (db.Notification, error)
	GetNotification(ctx context.Context, id int64) (db.Notification, error)
	ListNotifications(ctx context.Context, arg db.ListNotificationsParams) ([]db.Notification, error)
	UpdateNotification(ctx context.Context, arg db.UpdateNotificationParams) (db.Notification, error)
	DeleteNotification(ctx context.Context, id int64) error
}

type Repository struct {
	queries *db.Queries
}

func NewRepository(queries *db.Queries) RepositoryInterface {
	return &Repository{queries}
}

func (r *Repository) CreateNotification(ctx context.Context, arg db.CreateNotificationParams) (db.Notification, error) {
	return r.queries.CreateNotification(ctx, arg)
}

func (r *Repository) GetNotification(ctx context.Context, id int64) (db.Notification, error) {
	return r.queries.GetNotification(ctx, id)
}

func (r *Repository) ListNotifications(ctx context.Context, arg db.ListNotificationsParams) ([]db.Notification, error) {
	return r.queries.ListNotifications(ctx, arg)
}

func (r *Repository) UpdateNotification(ctx context.Context, arg db.UpdateNotificationParams) (db.Notification, error) {
	return r.queries.UpdateNotification(ctx, arg)
}

func (r *Repository) DeleteNotification(ctx context.Context, id int64) error {
	return r.queries.DeleteNotification(ctx, id)
}
