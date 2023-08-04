package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"
)

type Timeline interface {
	// Fetch account which has specified username
	FindPublicTimeline(ctx context.Context, MaxID int64, SinceID int64, Limit int64) (*object.Timeline, error)
	FindHomeTimeline(ctx context.Context, MaxID int64, SinceID int64, Limit int64, account *object.Account) (*object.Timeline, error)
	// TODO: Add Other APIs
	// CreateStatus(ctx context.Context, status *object.Status) (*object.Status, error)
}
