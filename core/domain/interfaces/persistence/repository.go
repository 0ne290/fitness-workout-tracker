package persistence

import (
	"context"

	"github.com/0ne290/fitness-workout-tracker/core/domain/entities"
)

type Repository interface {
	AddAdmin(ctx context.Context, admin *entities.Admin)
	TryGetAdminByGuid(ctx context.Context, guid []byte) *entities.Admin
}
