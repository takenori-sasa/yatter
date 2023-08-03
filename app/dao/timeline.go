package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Account
	timeline struct {
		db *sqlx.DB
	}
)

// Create accout repository
func NewTimeline(db *sqlx.DB) repository.Timeline {
	return &timeline{db: db}
}

func clampInt64(value, min, max int64) int64 {
	if value < min {
		return min
	} else if value > max {
		return max
	}
	return value
}

func (r *timeline) FindPublicTimeline(ctx context.Context, MaxID, SinceID, Limit int64) (*object.Timeline, error) {
	MaxID = clampInt64(MaxID, 1, math.MaxInt64)
	SinceID = clampInt64(SinceID, 0, math.MaxInt64)
	Limit = clampInt64(Limit, 0, 80)

	entity := new(object.Timeline)
	rows, err := r.db.QueryxContext(ctx, "SELECT * FROM status WHERE id > ? AND id<? LIMIT ? ", SinceID, MaxID, Limit)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find statuses from db: %w", err)
	}

	for rows.Next() {
		var s object.Status
		err := rows.StructScan(&s)
		if err != nil {
			return nil, fmt.Errorf("failed to scan status from db: %w", err)
		}
		entity.Body = append(entity.Body, &s)
	}
	return entity, nil
}

func (r *timeline) FindHomeTimeline(ctx context.Context, MaxID, SinceID, Limit int64, account *object.Account) (*object.Timeline, error) {
	MaxID = clampInt64(MaxID, 1, math.MaxInt64)
	SinceID = clampInt64(SinceID, 0, math.MaxInt64)
	Limit = clampInt64(Limit, 0, 80)

	entity := new(object.Timeline)
	rows, err := r.db.QueryxContext(ctx, "SELECT * FROM status WHERE id > ? AND id<? AND account_id=? LIMIT ? ", SinceID, MaxID, account.ID, Limit)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find statuses from db: %w", err)
	}

	for rows.Next() {
		var s object.Status
		err := rows.StructScan(&s)
		if err != nil {
			return nil, fmt.Errorf("failed to scan status from db: %w", err)
		}
		entity.Body = append(entity.Body, &s)
	}
	return entity, nil
}
