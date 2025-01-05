package infrastructure

import "time"

type TimeProvider struct{}

func (*TimeProvider) UtcNow() time.Time {
	return time.Now().UTC()
}
