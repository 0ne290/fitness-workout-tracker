package interfaces

import "time"

type TimeProvider interface {
	UtcNow() time.Time
}
