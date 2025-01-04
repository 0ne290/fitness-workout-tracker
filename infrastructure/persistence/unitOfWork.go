package infrastructure

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type unitOfWork struct {
	pool        *pgxpool.Pool
	transaction pgx.Tx
}

func NewUnitOfWork(pool *pgxpool.Pool) *unitOfWork {
	return &unitOfWork{pool, nil}
}

func (unitOfWork *unitOfWork) Begin(ctx context.Context) *repository {
	if unitOfWork.transaction != nil {
		unitOfWork.transaction.Rollback(ctx)
		panic("persistence.UnitOfWork.Begin(): transaction is invalid")
	}

	var err error
	unitOfWork.transaction, err = unitOfWork.pool.Begin(ctx)

	if err != nil {
		panic("persistence.UnitOfWork.Begin(): pool.Begin() error. Detail: " + err.Error())
	}

	return NewRepository(unitOfWork.transaction)
}

func (unitOfWork *unitOfWork) Save(ctx context.Context) {
	if unitOfWork.transaction == nil {
		panic("persistence.UnitOfWork.Save(): transaction is invalid")
	}

	if err := unitOfWork.transaction.Commit(ctx); err != nil {
		panic("persistence.UnitOfWork.Save(): transaction.Commit() error. Detail: " + err.Error())
	}
}
