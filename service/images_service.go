package service

import (
	"context"
	"github.com/ferripradana/jwt-authentication/model/web"
)

type ImageService interface {
	Create(ctx context.Context, request web.ImageCreateRequest) []web.ImageResponse
	Delete(ctx context.Context, imageId string)
	FindById(ctx context.Context, imageId string) web.ImageResponse
}
