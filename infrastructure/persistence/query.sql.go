// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package persistence

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createAdmin = `-- name: CreateAdmin :exec
INSERT INTO admins (guid, createdAt, name, login, password) VALUES (
    $1, $2, $3, $4, $5
)
`

type CreateAdminParams struct {
	Guid      pgtype.UUID
	Createdat pgtype.Timestamp
	Name      string
	Login     string
	Password  string
}

func (q *Queries) CreateAdmin(ctx context.Context, arg CreateAdminParams) error {
	_, err := q.db.Exec(ctx, createAdmin,
		arg.Guid,
		arg.Createdat,
		arg.Name,
		arg.Login,
		arg.Password,
	)
	return err
}

const createGeneralExercise = `-- name: CreateGeneralExercise :exec
INSERT INTO generalExercises (guid, adminGuid, name, description) VALUES (
    $1, $2, $3, $4
)
`

type CreateGeneralExerciseParams struct {
	Guid        pgtype.UUID
	Adminguid   pgtype.UUID
	Name        string
	Description string
}

func (q *Queries) CreateGeneralExercise(ctx context.Context, arg CreateGeneralExerciseParams) error {
	_, err := q.db.Exec(ctx, createGeneralExercise,
		arg.Guid,
		arg.Adminguid,
		arg.Name,
		arg.Description,
	)
	return err
}

const createUser = `-- name: CreateUser :exec
INSERT INTO users (guid, createdAt, name, login, password) VALUES (
    $1, $2, $3, $4, $5
)
`

type CreateUserParams struct {
	Guid      pgtype.UUID
	Createdat pgtype.Timestamp
	Name      string
	Login     string
	Password  string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.Exec(ctx, createUser,
		arg.Guid,
		arg.Createdat,
		arg.Name,
		arg.Login,
		arg.Password,
	)
	return err
}

const createUserExercise = `-- name: CreateUserExercise :exec
INSERT INTO userExercises (guid, userGuid, name, description) VALUES (
    $1, $2, $3, $4
)
`

type CreateUserExerciseParams struct {
	Guid        pgtype.UUID
	Userguid    pgtype.UUID
	Name        string
	Description string
}

func (q *Queries) CreateUserExercise(ctx context.Context, arg CreateUserExerciseParams) error {
	_, err := q.db.Exec(ctx, createUserExercise,
		arg.Guid,
		arg.Userguid,
		arg.Name,
		arg.Description,
	)
	return err
}

const createWorkout = `-- name: CreateWorkout :exec
INSERT INTO workouts (guid, userGuid, scheduledFor, durationInMinutes, setsAndBreaks) VALUES (
    $1, $2, $3, $4, $5
)
`

type CreateWorkoutParams struct {
	Guid              pgtype.UUID
	Userguid          pgtype.UUID
	Scheduledfor      pgtype.Timestamp
	Durationinminutes float64
	Setsandbreaks     []byte
}

func (q *Queries) CreateWorkout(ctx context.Context, arg CreateWorkoutParams) error {
	_, err := q.db.Exec(ctx, createWorkout,
		arg.Guid,
		arg.Userguid,
		arg.Scheduledfor,
		arg.Durationinminutes,
		arg.Setsandbreaks,
	)
	return err
}