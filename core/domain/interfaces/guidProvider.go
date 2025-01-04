package interfaces

type GuidProvider interface {
	Random() [16]byte
}
