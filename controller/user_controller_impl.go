package controller

import (
	"github.com/ferripradana/jwt-authentication/helper"
	"github.com/ferripradana/jwt-authentication/model/web"
	"github.com/ferripradana/jwt-authentication/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{UserService: userService}
}

func (controller *UserControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userCreateRequest := web.UserCreateRequest{}
	helper.ReadFromRequestBody(request, &userCreateRequest)

	response := controller.UserService.Create(request.Context(), userCreateRequest)
	webResponse := web.Response{
		Status: "Ok",
		Data:   response,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userUpdateRequest := web.UserUpdateRequest{}
	helper.ReadFromRequestBody(request, &userUpdateRequest)

	userUpdateRequest.Id = params.ByName("user_id")

	response := controller.UserService.Update(request.Context(), userUpdateRequest)
	webResponse := web.Response{
		Status: "Ok",
		Data:   response,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("user_id")
	controller.UserService.Delete(request.Context(), userId)
	webResponse := web.Response{
		Status: "Ok",
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("user_id")
	response := controller.UserService.FindById(request.Context(), userId)
	webResponse := web.Response{
		Status: "OK",
		Data:   response,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	response := controller.UserService.FindAll(request.Context())
	webResponse := web.Response{
		Status: "OK",
		Data:   response,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
