package main

import (
	"context"
	"crypto/sha512"
	"fmt"

	"github.com/0ne290/fitness-workout-tracker/core/domain/entities"
	"github.com/0ne290/fitness-workout-tracker/infrastructure"
	"github.com/0ne290/fitness-workout-tracker/infrastructure/persistence"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	admin, err := entities.NewAdmin(sha512.New(), &infrastructure.TimeProvider{}, &infrastructure.GuidProvider{}, "One290", "AnyLogin", "AnyPassword1")

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
	unitOfWork := persistence.NewUnitOfWork(pool)
	repository := unitOfWork.Begin(context)
	repository.AddAdmin(context, admin)
	unitOfWork.Save(context)
}
