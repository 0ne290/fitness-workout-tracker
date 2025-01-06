package persistence

import (
	"context"

	interfacesPersistence "github.com/0ne290/fitness-workout-tracker/core/domain/interfaces/persistence"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UnitOfWork struct {
	repository *Repository
}

func NewUnitOfWork(ctx context.Context, pool *pgxpool.Pool) *UnitOfWork {
	transaction, err := pool.Begin(ctx)

	if err != nil {
		panic("infrastructure.NewUnitOfWork(): pool.Begin() error. Detail: " + err.Error())
	}

	return &UnitOfWork{newRepository(transaction)}
}

func (unitOfWork *UnitOfWork) Repository() interfacesPersistence.Repository {
	return unitOfWork.repository
}

func (unitOfWork *UnitOfWork) Save(ctx context.Context) {
	if err := unitOfWork.repository.transaction.Commit(ctx); err != nil {
		panic("infrastructure.UnitOfWork.Save(): transaction.Commit() error. Detail: " + err.Error())
	}
}

func (unitOfWork *UnitOfWork) Rollback(ctx context.Context) {
	if err := unitOfWork.repository.transaction.Rollback(ctx); err != nil {
		panic("infrastructure.UnitOfWork.Rollback(): transaction.Rollback() error. Detail: " + err.Error())
	}
}
