package utils

import (
	"github.com/ferripradana/jwt-authentication/exception"
	"github.com/ferripradana/jwt-authentication/helper"
	"github.com/ferripradana/jwt-authentication/model/web"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

func CreateToken(request web.TokenCreateRequest, value time.Duration) string {
	var jwtTokenSecret = []byte(os.Getenv("JWT_TOKEN_SECRET"))

	expiredTime := time.Now().Add(time.Minute * value)
	claims := &web.TokenClaims{
		UserId:    request.UserId,
		Email:     request.Email,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{
				expiredTime,
			},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtTokenSecret)
	helper.IfErrorPanic(err)
	return tokenStr
}

func ClaimTokens(refreshToken string) web.TokenClaims {
	var jwtTokenSecret = []byte(os.Getenv("JWT_TOKEN_SECRET"))
	claims := &web.TokenClaims{}

	token, err := jwt.ParseWithClaims(
		refreshToken,
		claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtTokenSecret, nil
		},
	)

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			panic(exception.NewUnauthorizedError(err.Error()))
		}
	}

	if token == nil || !token.Valid {
		panic(exception.NewUnauthorizedError(err.Error()))
	}
	return *claims
}
