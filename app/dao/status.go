package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Status
	status struct {
		db *sqlx.DB
	}
)

// NewStatus creates a new status repository
func NewStatus(db *sqlx.DB) repository.Status {
	return &status{db: db}
}

// FindStatus retrieves a status by its ID
func (r *status) FindStatus(ctx context.Context, statusID int64) (*object.Status, error) {
	entity := new(object.Status)
	err := r.db.QueryRowxContext(ctx, "SELECT status.*, account.id  AS `account.id`, account.username AS `account.username`,account.create_at AS `account.create_at` FROM status JOIN account ON account.id=status.account_id WHERE status.id = ?", statusID).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find status from db: %w", err)
	}
	return entity, nil
}

// CreateStatus creates a new status in the database
func (r *status) CreateStatus(ctx context.Context, status *object.Status) (*object.Status, error) {
	res, err := r.db.Exec(`INSERT INTO status (content, account_id) values(?, ?)`, status.Content, status.AccountID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to create status: %w", err)
	}
	statusID, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get statusID: %w", err)
	}
	return r.FindStatus(ctx, statusID)
}

// DeleteStatus deletes a status by its ID
func (r *status) DeleteStatus(ctx context.Context, statusID int64) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}
	_, err = tx.ExecContext(ctx, "DELETE FROM status WHERE id = ?", statusID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		tx.Rollback()
		return fmt.Errorf("failed to delete status: %w", err)
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	return nil
}
