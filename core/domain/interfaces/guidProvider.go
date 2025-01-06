package interfaces

type GuidProvider interface {
	Random() []byte
	String(guid []byte) string
}
