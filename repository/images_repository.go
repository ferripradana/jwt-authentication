package repository

import (
	"context"
	"database/sql"
	"github.com/ferripradana/jwt-authentication/model/domain"
)

type ImageRepository interface {
	Create(ctx context.Context, tx *sql.Tx, image domain.Images) domain.Images
	Delete(ctx context.Context, tx *sql.Tx, image domain.Images)
	FindById(ctx context.Context, tx *sql.Tx, imageId string) (domain.Images, error)
}
