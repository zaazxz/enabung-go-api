package repository

import (
	"context"
	"database/sql"

	"github.com/zaazxz/enabung-go-api/model/domain"
)

type UsersRepository interface {
	Save(ctx context.Context, tx *sql.Tx, users domain.Users) (domain.Users, error)
	Update(ctx context.Context, tx *sql.Tx, users domain.Users) (domain.Users, error)
	Delete(ctx context.Context, tx *sql.Tx, users domain.Users) (domain.Users, error)
	FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.Users, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Users, error)
}
