package infrastructure

import "time"

type timeProvider struct{}

func (*timeProvider) UtcNow() time.Time {
	return time.Now().UTC()
}
