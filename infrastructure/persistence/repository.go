package persistence

import (
	"context"

	"github.com/0ne290/fitness-workout-tracker/core/domain/entities"
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	transaction pgx.Tx
}

func newRepository(transaction pgx.Tx) *Repository {
	return &Repository{transaction}
}

func (repository *Repository) AddAdmin(ctx context.Context, admin *entities.Admin) {
	if _, err := repository.transaction.Exec(ctx, "INSERT INTO admins VALUES ($1, $2, $3, $4, $5)", admin.Guid, admin.CreatedAt, admin.Name, admin.Login, admin.Password); err != nil {
		panic("infrastructure.Repository.AddAdmin(): transaction.Exec() error. Detail :" + err.Error())
	}
}
