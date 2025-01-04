SELECT 'CREATE DATABASE fitness_workout_tracker' 
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'fitness_workout_tracker');

CREATE TABLE IF NOT EXISTS fitness_workout_tracker.admins (
    guid UUID PRIMARY KEY,
    createdAt TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    login TEXT NOT NULL,
    password TEXT NOT NULL,
    CONSTRAINT uq_admins_login UNIQUE (login),
    CONSTRAINT uq_admins_loginAndPassword UNIQUE (login, password)
);

CREATE TABLE IF NOT EXISTS fitness_workout_tracker.users (
    guid UUID PRIMARY KEY,
    createdAt TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    login TEXT NOT NULL,
    password TEXT NOT NULL,
    CONSTRAINT uq_users_login UNIQUE (login),
    CONSTRAINT uq_users_loginAndPassword UNIQUE (login, password)
);

CREATE TABLE IF NOT EXISTS fitness_workout_tracker.generalExercises (
    guid UUID PRIMARY KEY,
    adminGuid UUID NOT NULL REFERENCES admins(guid),
    name TEXT NOT NULL,
    description TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS fitness_workout_tracker.userExercises (
    guid UUID PRIMARY KEY,
    userGuid UUID NOT NULL REFERENCES users(guid),
    name TEXT NOT NULL,
    description TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS fitness_workout_tracker.workouts (
    guid UUID PRIMARY KEY,
    userGuid UUID NOT NULL REFERENCES users(guid),
    scheduledFor TIMESTAMP NOT NULL,
    durationInMinutes DOUBLE PRECISION NOT NULL,
    setsAndBreaks JSONB NOT NULL
);