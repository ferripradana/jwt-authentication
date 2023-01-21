package main

import (
	"fmt"
	"github.com/ferripradana/jwt-authentication/app"
	"github.com/ferripradana/jwt-authentication/controller"
	"github.com/ferripradana/jwt-authentication/helper"
	"github.com/ferripradana/jwt-authentication/middleware"
	"github.com/ferripradana/jwt-authentication/repository"
	"github.com/ferripradana/jwt-authentication/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"net/http"
)

func main() {
	fmt.Println("Start...")
	errEnv := godotenv.Load(".env")
	helper.IfErrorPanic(errEnv)
	validate := validator.New()
	db := app.NewDB()

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)
	router := app.NewRouter(userController)

	authMiddleware := middleware.NewAuthMiddleware(router)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}

	err := server.ListenAndServe()
	helper.IfErrorPanic(err)

}
