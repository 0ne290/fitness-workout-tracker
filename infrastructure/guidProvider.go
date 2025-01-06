package infrastructure

import "github.com/google/uuid"

type GuidProvider struct{}

func (*GuidProvider) Random() []byte {
	guid := uuid.New()
	byteArrayOfGuid := [16]byte(guid)
	return byteArrayOfGuid[:]
}

func (*GuidProvider) String(guid []byte) string {
	uuid, err := uuid.FromBytes(guid)

	if err != nil {
		panic("infrastructure.GuidProvider.String(): guid is invalid")
	}

	return uuid.String()
}
