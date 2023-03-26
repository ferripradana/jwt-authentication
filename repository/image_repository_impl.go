package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/ferripradana/jwt-authentication/helper"
	"github.com/ferripradana/jwt-authentication/model/domain"
)

type ImageRepositoryImpl struct {
}

func NewImageRepositoryImpl() ImageRepository {
	return &ImageRepositoryImpl{}
}

func (repository *ImageRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, image domain.Images) domain.Images {
	SQL := "INSERT INTO images(id, path, created_at, updated_at) VALUES(?,?,?,?)"
	_, err := tx.ExecContext(
		ctx,
		SQL,
		image.Id,
		image.Path,
		image.CreatedAt,
		image.UpdatedAt,
	)
	helper.IfErrorPanic(err)
	return image
}

func (repository *ImageRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, image domain.Images) {
	SQL := "DELETE FROM images WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, image.Id)
	helper.IfErrorPanic(err)
}

func (repository *ImageRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, imageId string) (domain.Images, error) {
	SQL := "SELECT id, path, created_at, updated_at FROM images WHERE id=?"
	rows, err := tx.QueryContext(
		ctx,
		SQL,
		imageId,
	)
	helper.IfErrorPanic(err)
	defer rows.Close()
	image := domain.Images{}
	if rows.Next() {
		err := rows.Scan(
			&image.Id,
			&image.Path,
			&image.CreatedAt,
			&image.UpdatedAt,
		)
		helper.IfErrorPanic(err)
		return image, nil
	}
	return image, errors.New("image not found")
}
