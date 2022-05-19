package service

import (
	"context"
	"database/sql"
	"fmt"

	"nidzamTest.com/exception"
	"nidzamTest.com/helper"
	"nidzamTest.com/model"
	"nidzamTest.com/repository"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
}

func NewUserService(userRepo repository.UserRepository, DB *sql.DB) UserService {
	return &UserServiceImpl{
		UserRepository: userRepo,
		DB:             DB,
	}
}
func (service *UserServiceImpl) VerifyCredential(ctx context.Context, username string, password string) (interface{}, string) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	res := service.UserRepository.VerifyCredential(ctx, tx, username, password)
	if len(res) == 1 {
		return true, res[0].Role
	}

	return false, ""
}
func (service *UserServiceImpl) Insert(ctx context.Context, user model.UserEntity, username string, role string) string {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	if role != "admin" {
		var objectMessage []exception.BadRequestError
		var message exception.BadRequestError
		message.Desc = "role anda bukan admin"
		message.DescGlob = "role anda bukan admin"
		message.FieldName = "role"
		objectMessage = append(objectMessage, message)
		panic(exception.NewBadRequestError(objectMessage))
	}

	insertData := service.UserRepository.Insert(ctx, tx, user, username)
	if insertData == "gagal" {
		getDetail := service.UserRepository.FindAll(ctx, tx, fmt.Sprintf("username='%s'", user.Username))
		if len(getDetail) > 0 {
			var objectMessage []exception.BadRequestError
			var message exception.BadRequestError
			message.Desc = "username sudah ada"
			message.DescGlob = "username sudah ada"
			message.FieldName = "username"
			objectMessage = append(objectMessage, message)
			panic(exception.NewBadRequestError(objectMessage))

		}
	}
	return insertData

}
func (service *UserServiceImpl) Update(ctx context.Context, user model.UserEntity, username string, role string) string {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	if role != "admin" {
		var objectMessage []exception.BadRequestError
		var message exception.BadRequestError
		message.Desc = "role anda bukan admin"
		message.DescGlob = "role anda bukan admin"
		message.FieldName = "role"
		objectMessage = append(objectMessage, message)
		panic(exception.NewBadRequestError(objectMessage))
	}
	cekAda := service.UserRepository.FindAll(ctx, tx, fmt.Sprintf(" rowId=%d", user.RowId))
	var objectMessage []exception.BadRequestError
	var message exception.BadRequestError
	if len(cekAda) == 0 {
		message.Desc = "data tidak ada"
		message.DescGlob = "data tidak ada"
		message.FieldName = "rowId"
		objectMessage = append(objectMessage, message)
		panic(exception.NewBadRequestError(objectMessage))
	}
	updateData := service.UserRepository.Update(ctx, tx, user, username)
	if updateData == "gagal" {
		getDetail := service.UserRepository.FindAll(ctx, tx, fmt.Sprintf("username='%s' and rowId!=%d", user.Username, user.RowId))
		if len(getDetail) > 0 {
			message.Desc = "username sudah ada"
			message.DescGlob = "username sudah ada"
			message.FieldName = "username"
			objectMessage = append(objectMessage, message)
			panic(exception.NewBadRequestError(objectMessage))

		}
	}
	return updateData
}
func (service *UserServiceImpl) Delete(ctx context.Context, username string, role string) string {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	var objectMessage []exception.BadRequestError
	var message exception.BadRequestError
	if role != "admin" {
		message.Desc = "role anda bukan admin"
		message.DescGlob = "role anda bukan admin"
		message.FieldName = "role"
		objectMessage = append(objectMessage, message)
		panic(exception.NewBadRequestError(objectMessage))
	} else if username == "admin" {
		message.Desc = "tidak boleh di hapus"
		message.DescGlob = "tidak boleh di hapus"
		message.FieldName = "username"
		objectMessage = append(objectMessage, message)
		panic(exception.NewBadRequestError(objectMessage))
	}
	deleteData := service.UserRepository.Delete(ctx, tx, username)
	return deleteData
}
func (service *UserServiceImpl) FindByUsername(ctx context.Context, username string) interface{} {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	getData := service.UserRepository.FindByUsername(ctx, tx, username)
	if len(getData) > 0 {
		return getData[0]
	} else {
		return nil
	}
}
func (service *UserServiceImpl) FindAll(ctx context.Context, page int, perpage int, filter string, order string) []model.UserEntity {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	if filter == "" {
		filter = "1=1"
	}
	if order == "" {
		order = "username asc"
	}
	if page > 0 {
		page = page - 1
	}
	where := fmt.Sprintf("%s order by %s OFFSET %d ROWS FETCH NEXT %d ROWS ONLY", filter, order, page, perpage)
	getData := service.UserRepository.FindAll(ctx, tx, where)
	return getData
}
