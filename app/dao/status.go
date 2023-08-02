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
	// Implementation for repository.Account
	status struct {
		db *sqlx.DB
	}
)

// Create accout repository
func NewStatus(db *sqlx.DB) repository.Status {
	return &status{db: db}
}

// FindByUsername : ユーザ名からユーザを取得
func (r *status) FindStatus(ctx context.Context, status *object.Status) (*object.Status, error) {
	entity := new(object.Account)
	err := r.db.QueryRowxContext(ctx, "select * from account where username = ?", username).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to find account from db: %w", err)
	}

	return entity, nil
}
func (r *status) CreateStatus(ctx context.Context, status *object.Status) (*object.Status, error) {
	_, err := r.db.Exec(`INSERT INTO status (username, password_hash) values(?, ?)`, account.Username, account.PasswordHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to create account: %w", err)
	}

	return r.FindByUsername(ctx, account.Username)
}
