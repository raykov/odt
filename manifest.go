package odt

type Manifest struct{}

const manifestFile = "manifest.rdf"

func (m *Manifest) Write(w zipWriter) error {
	f, err := w.Create(manifestFile)
	if err != nil {
		return err
	}

	_, err = f.Write(m.manifest())
	if err != nil {
		return err
	}

	return nil
}

func (m *Manifest) manifest() []byte {
	return []byte(`<?xml version="1.0" encoding="utf-8"?>
<rdf:RDF xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#">
  <rdf:Description rdf:about="">
    <rdf:type rdf:resource="http://docs.oasis-open.org/ns/office/1.2/meta/pkg#Document"/>
  </rdf:Description>
</rdf:RDF>`)
}
