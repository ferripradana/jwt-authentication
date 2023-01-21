package app

import (
	"github.com/ferripradana/jwt-authentication/controller"
	"github.com/ferripradana/jwt-authentication/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(userController controller.UserController) *httprouter.Router {
	router := httprouter.New()
	router.GET("/api/v1/users", userController.FindAll)
	router.GET("/api/v1/users/:user_id", userController.FindById)
	router.POST("/api/v1/users", userController.Create)
	router.PUT("/api/v1/users/:user_id", userController.Update)
	router.DELETE("/api/v1/users/:user_id", userController.Delete)
	router.POST("/api/v1/auth", userController.Auth)
	router.POST("/api/v1/refresh-token", userController.CreateWithRefreshToken)

	router.PanicHandler = exception.ErrorHandler

	return router
}
