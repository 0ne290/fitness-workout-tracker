package persistence

import "context"

type UnitOfWork interface {
	Repository() Repository
	Save(ctx context.Context)
	Rollback(ctx context.Context)
}
