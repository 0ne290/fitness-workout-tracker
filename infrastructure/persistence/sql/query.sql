-- CreateAdmin :exec
INSERT INTO admins VALUES (
    $1, $2, $3, $4, $5
);

-- CreateUser :exec
INSERT INTO users VALUES (
    $1, $2, $3, $4, $5
);

-- CreateGeneralExercise :exec
INSERT INTO generalExercises VALUES (
    $1, $2, $3, $4
);

-- CreateUserExercise :exec
INSERT INTO userExercises VALUES (
    $1, $2, $3, $4
);

-- CreateWorkout :exec
INSERT INTO workouts VALUES (
    $1, $2, $3, $4, $5
);