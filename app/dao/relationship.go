package dao

import (
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Account
	relationship struct {
		db *sqlx.DB
	}
)

// NewAccount creates a new account repository
func NewRelationship(db *sqlx.DB) repository.Relationship {
	return &account{db: db}
}
