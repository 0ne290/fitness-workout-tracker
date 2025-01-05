package persistence

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UnitOfWork struct {
	pool        *pgxpool.Pool
	transaction pgx.Tx
}

func NewUnitOfWork(pool *pgxpool.Pool) *UnitOfWork {
	return &UnitOfWork{pool, nil}
}

func (unitOfWork *UnitOfWork) Begin(ctx context.Context) *Repository {
	var err error
	unitOfWork.transaction, err = unitOfWork.pool.Begin(ctx)

	if err != nil {
		panic("infrastructure.UnitOfWork.Begin(): pool.Begin() error. Detail: " + err.Error())
	}

	return newRepository(unitOfWork.transaction)
}

func (unitOfWork *UnitOfWork) Save(ctx context.Context) {
	if err := unitOfWork.transaction.Commit(ctx); err != nil {
		panic("infrastructure.UnitOfWork.Save(): transaction.Commit() error. Detail: " + err.Error())
	}
}
