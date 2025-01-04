package infrastructure

import (
	"context"

	"github.com/0ne290/fitness-workout-tracker/core/domain/entities"
	"github.com/jackc/pgx/v5"
)

type repository struct {
	transaction pgx.Tx
}

func NewRepository(transaction pgx.Tx) *repository {
	return &repository{transaction}
}

func (repository *repository) AddAdmin(ctx context.Context, admin entities.Admin) {
	repository.transaction.Exec(ctx, "INSERT INTO admins VALUES ($1, $2, $3, $4, $5)", admin.Guid, admin.CreatedAt, admin.Name, admin.Login, admin.Password)
}
