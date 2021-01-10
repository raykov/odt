package odt

type DocElement interface {
	Write() []byte
	String() string
}
