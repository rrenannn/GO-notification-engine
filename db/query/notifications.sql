-- name: CreateNotification :one
INSERT INTO notifications (recipient, message, channel_type, status)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetNotification :one
SELECT * FROM notifications
WHERE id = $1 LIMIT 1;

-- name: ListNotifications :many
SELECT * FROM notifications
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateNotification :exec
UPDATE notifications
SET status = $2,
    updated_at = NOW()
WHERE id = $1;

-- name: DeleteNotification :exec
DELETE FROM notifications
WHERE id = $1;