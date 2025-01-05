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

	pool, err := pgxpool.New(context.Background(), "User ID=root;Password=dcba;Host=localhost;Port=5432;Database=fitness-workout-tracker;Pooling=true;Min Pool Size=0;Max Pool Size=100;Connection Lifetime=0;")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	context := context.Background()
	unitOfWork := persistence.NewUnitOfWork(pool)
	repository := unitOfWork.Begin(context)
	repository.AddAdmin(context, admin)
}
