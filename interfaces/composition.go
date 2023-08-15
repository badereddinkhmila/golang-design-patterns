package interfaces

type Writer interface {
	Write(data []byte) (int, error)
}

type Reader interface {
	Read() ([]byte, error)
}

type ReadWriter interface {
	Reader
	Writer
}
