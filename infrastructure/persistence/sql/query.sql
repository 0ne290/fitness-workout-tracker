-- name: CreateAdmin :exec
INSERT INTO admins VALUES (
    $1, $2, $3, $4, $5
);

-- name: CreateUser :exec
INSERT INTO users VALUES (
    $1, $2, $3, $4, $5
);

-- name: CreateGeneralExercise :exec
INSERT INTO generalExercises VALUES (
    $1, $2, $3, $4
);

-- name: CreateUserExercise :exec
INSERT INTO userExercises VALUES (
    $1, $2, $3, $4
);

-- name: CreateWorkout :exec
INSERT INTO workouts VALUES (
    $1, $2, $3, $4, $5
);