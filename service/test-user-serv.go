package service

import (
	"context"

	"nidzamTest.com/model"
)

type UserService interface {
	VerifyCredential(ctx context.Context, username string, password string) (interface{}, string)
	Insert(ctx context.Context, user model.UserEntity, username string, role string) string
	Update(ctx context.Context, user model.UserEntity, username string, role string) string
	Delete(ctx context.Context, username string, role string) string
	FindByUsername(ctx context.Context, username string) interface{}
	FindAll(ctx context.Context, page int, perpage int, filter string, order string) []model.UserEntity
}
