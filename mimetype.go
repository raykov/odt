package odt

type Mimetype struct{}

const mimetypeFile = "mimetype"

func (m *Mimetype) Write(w zipWriter) error {
	f, err := w.Create(mimetypeFile)
	if err != nil {
		return err
	}

	_, err = f.Write([]byte("application/vnd.oasis.opendocument.text"))
	if err != nil {
		return err
	}

	return nil
}
