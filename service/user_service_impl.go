package service

import (
	"context"
	"database/sql"
	"github.com/ferripradana/jwt-authentication/exception"
	"github.com/ferripradana/jwt-authentication/helper"
	"github.com/ferripradana/jwt-authentication/model/domain"
	"github.com/ferripradana/jwt-authentication/model/web"
	"github.com/ferripradana/jwt-authentication/repository"
	"github.com/ferripradana/jwt-authentication/utils"
	"github.com/go-playground/validator/v10"
	"time"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, db *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             db,
		Validate:       validate,
	}
}

func (service *UserServiceImpl) Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.IfErrorPanic(err)

	tx, err := service.DB.Begin()
	helper.IfErrorPanic(err)
	defer helper.CommitOrRollback(tx)

	hashedPassword, err := utils.HashPassword(request.Password)
	helper.IfErrorPanic(err)

	user := domain.User{
		Id:        utils.Uuid(),
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	user = service.UserRepository.Create(ctx, tx, user)
	return utils.ToUserResponse(user)
}

func (service *UserServiceImpl) Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.IfErrorPanic(err)

	tx, err := service.DB.Begin()
	helper.IfErrorPanic(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	user.FirstName = request.FirstName
	user.LastName = request.LastName
	user.UpdatedAt = time.Now().Unix()

	user = service.UserRepository.Update(ctx, tx, user)
	return utils.ToUserResponse(user)
}

func (service *UserServiceImpl) Delete(ctx context.Context, user_id string) {
	tx, err := service.DB.Begin()
	helper.IfErrorPanic(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, user_id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.UserRepository.Delete(ctx, tx, user)
}

func (service *UserServiceImpl) FindById(ctx context.Context, user_id string) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.IfErrorPanic(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, user_id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return utils.ToUserResponse(user)
}

func (service *UserServiceImpl) FindAll(ctx context.Context) []web.UserResponse {
	tx, err := service.DB.Begin()
	helper.IfErrorPanic(err)
	defer helper.CommitOrRollback(tx)

	users := service.UserRepository.FindAll(ctx, tx)
	return utils.ToUserResponses(users)
}
