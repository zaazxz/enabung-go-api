package repository

import (
	"context"
	"database/sql"

	"github.com/zaazxz/enabung-go-api/helper"
	"github.com/zaazxz/enabung-go-api/model/domain"
)

type UsersRepositoryImpl struct{}

// Delete implements UsersRepository.
func (repository *UsersRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, users domain.Users) (domain.Users, error) {
	panic("unimplemented")
}

// FindAll implements UsersRepository.
func (repository *UsersRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Users, error) {
	panic("unimplemented")
}

// FindById implements UsersRepository.
func (repository *UsersRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.Users, error) {
	panic("unimplemented")
}

// Save implements UsersRepository.
func (repository *UsersRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, users domain.Users) (domain.Users, error) {
	query := "INSERT INTO users (name, email, balance, password) VALUES (?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, query, users.Name, users.Email, users.Balance, users.Password)
	if err != nil {
		return users, err
	}

	id, err := result.LastInsertId()
	helper.PanicError(err)
}

// Update implements UsersRepository.
func (repository *UsersRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, users domain.Users) (domain.Users, error) {
	panic("unimplemented")
}

func NewUsersRepository() UsersRepository {
	return &UsersRepositoryImpl{}
}
