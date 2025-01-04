package infrastructure

import "github.com/google/uuid"

type GuidProvider struct{}

func (*GuidProvider) Random() [16]byte {
	return [16]byte(uuid.New())
}
