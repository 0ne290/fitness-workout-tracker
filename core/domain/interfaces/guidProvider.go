package interfaces

type GuidProvider interface {
	Random() []byte
}
