package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/ferripradana/jwt-authentication/helper"
	"github.com/ferripradana/jwt-authentication/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (userRepositoryImpl *UserRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "insert into users(id, first_name, last_name, email, password, created_at, updated_at)" +
		"values(?,?,?,?,?,?,?)"
	_, err := tx.ExecContext(
		ctx,
		SQL,
		user.Id,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt,
	)
	helper.IfErrorPanic(err)
	return user
}

func (userRepositoryImpl *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "update users set first_name = ? , last_name = ? , updated_at = ? where id = ?"
	_, err := tx.ExecContext(
		ctx,
		SQL,
		user.FirstName,
		user.LastName,
		user.UpdatedAt,
		user.Id,
	)
	helper.IfErrorPanic(err)
	return user
}

func (userRepositoryImpl *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user domain.User) {
	SQL := "delete from users where id = ?"
	_, err := tx.ExecContext(
		ctx,
		SQL,
		user.Id,
	)
	helper.IfErrorPanic(err)
}

func (userRepositoryImpl *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, user_id string) (domain.User, error) {
	SQL := "select id, first_name, last_name, email, created_at, updated_at from users where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, user_id)
	helper.IfErrorPanic(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		helper.IfErrorPanic(err)
		return user, nil
	} else {
		return user, errors.New("User Not Found")
	}
}

func (userRepositoryImpl *UserRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error) {
	SQL := "select id, first_name, last_name, email, created_at, updated_at from users where email = ?"
	rows, err := tx.QueryContext(ctx, SQL, email)
	helper.IfErrorPanic(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		helper.IfErrorPanic(err)
		return user, nil
	} else {
		return user, errors.New("User Not Found")
	}
}

func (userRepositoryImpl *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	SQL := "select id, first_name, last_name, email, created_at, updated_at from users"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.IfErrorPanic(err)

	var users []domain.User
	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		helper.IfErrorPanic(err)
		users = append(users, user)
	}
	return users
}
