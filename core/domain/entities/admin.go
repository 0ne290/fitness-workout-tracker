package entities

import (
	"errors"
	"hash"
	"strings"
	"time"

	"github.com/0ne290/fitness-workout-tracker/core/domain/interfaces"
)

type Admin struct {
	Guid      []byte
	CreatedAt time.Time
	Name      string
	Login     string
	Password  []byte
}

func NewAdmin(hasher hash.Hash, timeProvider interfaces.TimeProvider, guidProvider interfaces.GuidProvider, name string, login string, password string) (*Admin, error) {
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		return nil, errors.New("name is invalid")
	}

	login = strings.TrimSpace(login)
	if len(login) < 6 {
		return nil, errors.New("login is invalid")
	}

	password = strings.TrimSpace(password)
	if len(password) < 12 {
		return nil, errors.New("password is invalid")
	}

	guid := guidProvider.Random()
	createdAt := timeProvider.UtcNow()
	hasher.Write(Salt(guid, password, createdAt))

	return &Admin{guid, createdAt, name, login, hasher.Sum(nil)}, nil
}

func Salt(guid []byte, password string, createdAt time.Time) []byte {
	return []byte(string(guid) + password + createdAt.String() + "|oIFK>w static_part_of_salt ;wI+Jrg")
}
