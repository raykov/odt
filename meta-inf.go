package odt

type MetaInf struct{}

const metaInfFile = "META-INF/manifest.xml"

func (m *MetaInf) Write(w zipWriter) error {
	f, err := w.Create(metaInfFile)
	if err != nil {
		return err
	}

	_, err = f.Write(m.manifest())
	if err != nil {
		return err
	}

	return nil
}

func (m *MetaInf) manifest() []byte {
	return []byte(`<?xml version="1.0" encoding="UTF-8"?>
<manifest:manifest xmlns:manifest="urn:oasis:names:tc:opendocument:xmlns:manifest:1.0" manifest:version="1.2" xmlns:loext="urn:org:documentfoundation:names:experimental:office:xmlns:loext:1.0">
 <manifest:file-entry manifest:full-path="/" manifest:version="1.2" manifest:media-type="application/vnd.oasis.opendocument.text"/>
 <manifest:file-entry manifest:full-path="content.xml" manifest:media-type="text/xml"/>
 <manifest:file-entry manifest:full-path="styles.xml" manifest:media-type="text/xml"/>
 <manifest:file-entry manifest:full-path="settings.xml" manifest:media-type="text/xml"/>
 <manifest:file-entry manifest:full-path="meta.xml" manifest:media-type="text/xml"/>
 <manifest:file-entry manifest:full-path="manifest.rdf" manifest:media-type="application/rdf+xml"/>
</manifest:manifest>`)
}
