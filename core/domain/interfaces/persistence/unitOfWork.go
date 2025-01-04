package interfaces

import "context"

type unitOfWork interface {
	Begin(ctx context.Context) *repository
	Save(ctx context.Context)
}
