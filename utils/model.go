package utils

import (
	"github.com/ferripradana/jwt-authentication/model/domain"
	"github.com/ferripradana/jwt-authentication/model/web"
)

func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToUserResponses(users []domain.User) []web.UserResponse {
	var userList []web.UserResponse
	for _, user := range users {
		userList = append(userList, ToUserResponse(user))
	}
	return userList
}
