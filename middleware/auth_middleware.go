package middleware

import (
	"github.com/ferripradana/jwt-authentication/helper"
	"github.com/ferripradana/jwt-authentication/model/web"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"os"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (authMiddleware *AuthMiddleware) unauthorized(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusUnauthorized)
	webResponse := web.Response{
		Status: "Unauthorized",
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (authMiddleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost && (request.RequestURI == "/api/v1/users" || request.RequestURI == "/api/v1/auth") {
		authMiddleware.Handler.ServeHTTP(writer, request)
	} else {
		tokenAuth := request.Header.Get("Authorization")
		if tokenAuth == "" {
			authMiddleware.unauthorized(writer, request)
			return
		}

		var jwtTokenSecret = []byte(os.Getenv("JWT_TOKEN_SECRET"))
		claims := &web.TokenClaims{}

		token, err := jwt.ParseWithClaims(
			tokenAuth,
			claims,
			func(t *jwt.Token) (interface{}, error) {
				return jwtTokenSecret, nil
			},
		)

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				authMiddleware.unauthorized(writer, request)
				return
			}
		}

		if token == nil || !token.Valid {
			authMiddleware.unauthorized(writer, request)
			return
		}

		authMiddleware.Handler.ServeHTTP(writer, request)
	}
}
