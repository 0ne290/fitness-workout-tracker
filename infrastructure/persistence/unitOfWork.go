package unitOfWork

type UnitOfWork struct {
    pool *pgx.Pool,
    transaction *pgxpool.Tx
}

func New(pool pgxpool.Pool) *UnitOfWork {
    return UnitOfWork{&pool, nil}
}

func (unitOfWork *UnitOfWork) Begin() Repository {
    if unitOfWork.transaction != nil {
        unitOfWork.transaction.Rollback()
        panic("unitOfWork.UnitOfWork.Begin(): transaction is invalid.")
    }

    unitOfWork.transaction = &pool.Begin()
    return Repository{unitOfWork.transaction}
}

func (unitOfWork *UnitOfWork) Save() {
    if unitOfWork.transaction == nil {
        panic("unitOfWork.UnitOfWork.Save(): transaction is invalid.")
    }

    if err := unitOfWork.transaction.Commit(); err != nil {
        panic("unitOfWork.UnitOfWork.Save(): transaction.Commit() error. Detail: {err}")
    }
}
