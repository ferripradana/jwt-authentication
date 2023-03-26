package controller

import (
	"github.com/ferripradana/jwt-authentication/helper"
	"github.com/ferripradana/jwt-authentication/model/web"
	"github.com/ferripradana/jwt-authentication/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
)

type ImageControllerImpl struct {
	ImageService service.ImageService
}

func NewImageControllerImpl(imageService service.ImageService) ImageController {
	return &ImageControllerImpl{ImageService: imageService}
}

func (controller *ImageControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	request.ParseMultipartForm(10 * 1024 * 1024)
	imageCreateRequest := web.ImageCreateRequest{}
	imageCreateRequest.FormData = request.MultipartForm.File["image"]
	imageCreated := controller.ImageService.Create(request.Context(), imageCreateRequest)
	response := web.Response{
		Status: "OK",
		Data:   imageCreated,
	}
	helper.WriteToResponseBody(writer, response)
}

func (controller *ImageControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	imageId := params.ByName("imageId")
	controller.ImageService.Delete(request.Context(), imageId)
	response := web.Response{
		Status: "OK",
	}
	helper.WriteToResponseBody(writer, response)
}

func (controller *ImageControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	imageId := params.ByName("imageId")
	image := controller.ImageService.FindById(request.Context(), imageId)
	fileBytes, err := os.ReadFile("public/" + image.Path)
	helper.IfErrorPanic(err)
	writer.Write(fileBytes)
}
