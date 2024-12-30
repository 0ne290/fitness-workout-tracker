-- name: CreateAdmin :exec
INSERT INTO admins (guid, createdAt, name, login, password) VALUES (
    $1, $2, $3, $4, $5
);

-- name: CreateUser :exec
INSERT INTO users (guid, createdAt, name, login, password) VALUES (
    $1, $2, $3, $4, $5
);

-- name: CreateGeneralExercise :exec
INSERT INTO generalExercises (guid, adminGuid, name, description) VALUES (
    $1, $2, $3, $4
);

-- name: CreateUserExercise :exec
INSERT INTO userExercises (guid, userGuid, name, description) VALUES (
    $1, $2, $3, $4
);

-- name: CreateWorkout :exec
INSERT INTO workouts (guid, userGuid, scheduledFor, durationInMinutes, setsAndBreaks) VALUES (
    $1, $2, $3, $4, $5
);