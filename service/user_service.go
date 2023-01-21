package service

import (
	"context"
	"github.com/ferripradana/jwt-authentication/model/web"
)

type UserService interface {
	Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse
	Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse
	Delete(ctx context.Context, user_id string)
	FindById(ctx context.Context, user_id string) web.UserResponse
	FindAll(ctx context.Context) []web.UserResponse
}
