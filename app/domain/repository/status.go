package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"
)

type Status interface {
	// Fetch account which has specified username
	FindStatus(ctx context.Context, statusID int64) (*object.Status, error)
	// TODO: Add Other APIs
	CreateStatus(ctx context.Context, status *object.Status) (*object.Status, error)
}
