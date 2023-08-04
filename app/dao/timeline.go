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
	// timeline is an implementation for repository.Timeline.
	// It provides methods to interact with timeline data in the database.
	timeline struct {
		db *sqlx.DB
	}
)

// NewTimeline creates a new repository.Timeline instance.
// It returns this instance which can be used to interact with timeline data in the database.
func NewTimeline(db *sqlx.DB) repository.Timeline {
	return &timeline{db: db}
}

func (r *timeline) FindPublicTimeline(ctx context.Context, MaxID int64, SinceID int64, Limit int64) (*object.Timeline, error) {
	// Input parameters validation
	if MaxID <= 0 {
		MaxID = math.MaxInt64
	}
	if SinceID < 0 {
		SinceID = 0
	}
	if Limit < 0 {
		Limit = 40
	}
	if Limit > 80 {
		Limit = 80
	}

	entity := new(object.Timeline)

	// Query to fetch the public timeline statuses
	rows, err := r.db.QueryxContext(ctx, "SELECT * FROM status WHERE id > ? AND id < ? LIMIT ?", SinceID, MaxID, Limit)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// No rows returned, return nil timeline
			return nil, nil
		}
		// Error occurred during query execution
		return nil, fmt.Errorf("failed to find statuses from db: %w", err)
	}

	for rows.Next() {
		var s object.Status

		// Scanning row into status object
		err := rows.StructScan(&s)
		if err != nil {
			// Error occurred during row scanning
			return nil, fmt.Errorf("failed to scan status from db: %w", err)
		}

		// Append scanned status to the timeline
		entity.Body = append(entity.Body, &s)
	}

	// Return fetched timeline
	return entity, nil
}

func (r *timeline) FindHomeTimeline(ctx context.Context, MaxID int64, SinceID int64, Limit int64, account *object.Account) (*object.Timeline, error) {
	// Input parameters validation
	if MaxID <= 0 {
		MaxID = math.MaxInt64
	}
	if SinceID < 0 {
		SinceID = 0
	}
	if Limit < 0 {
		Limit = 40
	}
	if Limit > 80 {
		Limit = 80
	}

	entity := new(object.Timeline)

	// Query to fetch the home timeline statuses for specific account
	rows, err := r.db.QueryxContext(ctx, "SELECT * FROM status WHERE id > ? AND id < ? AND account_id = ? LIMIT ?", SinceID, MaxID, account.ID, Limit)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// No rows returned, return nil timeline
			return nil, nil
		}
		// Error occurred during query execution
		return nil, fmt.Errorf("failed to find statuses from db: %w", err)
	}

	for rows.Next() {
		var s object.Status

		// Scanning row into status object
		err := rows.StructScan(&s)
		if err != nil {
			// Error occurred during row scanning
			return nil, fmt.Errorf("failed to scan status from db: %w", err)
		}

		// Append scanned status to the timeline
		entity.Body = append(entity.Body, &s)
	}

	// Return fetched timeline
	return entity, nil
}
