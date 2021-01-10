package odt

import "io"

type zipWriter interface {
	//SetOffset(n int64)
	//Flush() error
	//SetComment(comment string) error
	Close() error
	Create(name string) (io.Writer, error)
	//CreateHeader(fh *FileHeader) (io.Writer, error)
	//RegisterCompressor(method uint16, comp Compressor)
}
