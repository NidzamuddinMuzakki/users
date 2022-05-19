package repository

import (
	"context"
	"database/sql"

	"nidzamTest.com/model"
)

type UserRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, user model.UserEntity, username string) string
	Update(ctx context.Context, tx *sql.Tx, user model.UserEntity, username string) string
	Delete(ctx context.Context, tx *sql.Tx, username string) string
	FindByUsername(ctx context.Context, tx *sql.Tx, username string) []model.UserEntity

	FindAll(ctx context.Context, tx *sql.Tx, where string) []model.UserEntity
	VerifyCredential(ctx context.Context, tx *sql.Tx, username string, password string) []model.UserEntity
}
