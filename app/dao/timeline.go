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

func (r *timeline) FindPublicTimeline(ctx context.Context, MaxID int64, SinceID int64, Limit int64) (*object.Timeline, error) {
	if MaxID <= 0 {
		MaxID = math.MaxInt64
	}
	if SinceID < 0 {
		SinceID = 0
	}
	if Limit < 0 {
		Limit = 40
	}
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
		// fmt.Fprintf(s)
		entity.Body = append(entity.Body, &s)
	}
	return entity, nil

}

// // FindByUsername : ユーザ名からユーザを取得
// func (r *status) FindStatus(ctx context.Context, statusID int64) (*object.Status, error) {
// 	entity := new(object.Status)
// 	// err := r.db.QueryRowxContext(ctx, "SELECT * FROM status WHERE status.id = ?", statusID).StructScan(entity)
// 	// err := r.db.QueryRowxContext(ctx, "SELECT status.* FROM status JOIN account ON account.id=status.account_id WHERE status.id = ?", statusID).StructScan(entity)
// 	err := r.db.QueryRowxContext(ctx, "SELECT status.*, account.id  AS `account.id`, account.username AS `account.username` FROM status JOIN account ON account.id=status.account_id WHERE status.id = ?", statusID).StructScan(entity)
// 	if err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			return nil, nil
// 		}
//
// 		return nil, fmt.Errorf("failed to find status from db: %w", err)
// 	}
//
// 	return entity, nil
// }
//
// func (r *status) CreateStatus(ctx context.Context, status *object.Status) (*object.Status, error) {
// 	// create
// 	res, err := r.db.Exec(`INSERT INTO status (content, account_id) values(?, ?)`, status.Content, status.AccountID)
// 	if err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			return nil, nil
// 		}
//
// 		return nil, fmt.Errorf("failed to create status: %w", err)
// 	}
// 	// primary取得
// 	statusID, err := res.LastInsertId()
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to get statusID: %w", err)
// 	}
//
// 	return r.FindStatus(ctx, statusID)
// }
