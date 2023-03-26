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
	"io"
	"os"
	"strings"
	"time"
)

type ImageServiceImpl struct {
	ImageRepository repository.ImageRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewImageServiceImpl(imageRepository repository.ImageRepository, db *sql.DB, validate *validator.Validate) ImageService {
	return &ImageServiceImpl{
		ImageRepository: imageRepository,
		DB:              db,
		Validate:        validate,
	}
}

func (service *ImageServiceImpl) Create(ctx context.Context, request web.ImageCreateRequest) []web.ImageResponse {
	err := service.Validate.Struct(request)
	helper.IfErrorPanic(err)

	tx, err := service.DB.Begin()
	helper.IfErrorPanic(err)
	defer helper.CommitOrRollback(tx)

	var imageResponses []web.ImageResponse

	for _, image := range request.FormData {
		file, _ := image.Open()
		tempFile, err := os.CreateTemp("public", "images-*.jpg")
		helper.IfErrorPanic(err)
		defer tempFile.Close()

		fileBytes, err := io.ReadAll(file)
		helper.IfErrorPanic(err)

		tempFile.Write(fileBytes)
		fileName := tempFile.Name()
		newFileName := strings.Split(fileName, "\\")
		image := domain.Images{
			Id:        utils.Uuid(),
			Path:      newFileName[1],
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		}

		image = service.ImageRepository.Create(ctx, tx, image)
		imageResponses = append(imageResponses, utils.ToImageResponse(image))
	}

	return imageResponses
}

func (service *ImageServiceImpl) Delete(ctx context.Context, imageId string) {
	tx, err := service.DB.Begin()
	helper.IfErrorPanic(err)
	defer helper.CommitOrRollback(tx)

	image, err := service.ImageRepository.FindById(ctx, tx, imageId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.ImageRepository.Delete(ctx, tx, image)
	os.Remove("public/" + image.Path)
}

func (service *ImageServiceImpl) FindById(ctx context.Context, imageId string) web.ImageResponse {
	tx, err := service.DB.Begin()
	helper.IfErrorPanic(err)
	defer helper.CommitOrRollback(tx)

	image, err := service.ImageRepository.FindById(ctx, tx, imageId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return utils.ToImageResponse(image)
}
