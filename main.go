package main

import (
	"fmt"
	"github.com/ferripradana/jwt-authentication/app"
	"github.com/ferripradana/jwt-authentication/controller"
	"github.com/ferripradana/jwt-authentication/helper"
	"github.com/ferripradana/jwt-authentication/repository"
	"github.com/ferripradana/jwt-authentication/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {
	fmt.Println("Start...")
	validate := validator.New()
	db := app.NewDB()

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)
	router := app.NewRouter(userController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.IfErrorPanic(err)

}
