package main

import (
	"context"
	"crypto/sha512"
	"encoding/hex"
	"fmt"

	"github.com/0ne290/fitness-workout-tracker/core/domain/entities"
	interfacesPersistence "github.com/0ne290/fitness-workout-tracker/core/domain/interfaces/persistence"
	"github.com/0ne290/fitness-workout-tracker/infrastructure"
	infrastructurePersistence "github.com/0ne290/fitness-workout-tracker/infrastructure/persistence"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	guidProvider := &infrastructure.GuidProvider{}
	admin, err := entities.NewAdmin(sha512.New(), &infrastructure.TimeProvider{}, guidProvider, "One290", "AnyLogin", "AnyPassword1")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	pool, err := pgxpool.New(context.Background(), "user=root password=dcba host=localhost port=5432 dbname=fitness_workout_tracker")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	context := context.Background()
	var unitOfWork interfacesPersistence.UnitOfWork

	unitOfWork = infrastructurePersistence.NewUnitOfWork(context, pool)
	unitOfWork.Repository().AddAdmin(context, admin)
	unitOfWork.Save(context)

	unitOfWork = infrastructurePersistence.NewUnitOfWork(context, pool)
	admin1 := unitOfWork.Repository().TryGetAdminByGuid(context, admin.Guid)
	unitOfWork.Rollback(context)

	fmt.Println("Admin:")
	fmt.Println(guidProvider.String(admin.Guid))
	fmt.Println(admin.CreatedAt)
	fmt.Println(admin.Name)
	fmt.Println(admin.Login)
	fmt.Println(hex.EncodeToString(admin.Password))

	fmt.Println("\nAdmin1:")
	fmt.Println(guidProvider.String(admin1.Guid))
	fmt.Println(admin1.CreatedAt)
	fmt.Println(admin1.Name)
	fmt.Println(admin1.Login)
	fmt.Println(hex.EncodeToString(admin1.Password))
}
