package infrastructure

import "github.com/google/uuid"

type GuidProvider struct{}

func (*GuidProvider) Random() []byte {
	guid := uuid.New()
	byteArrayOfGuid := [16]byte(guid)
	return byteArrayOfGuid[:]
}
