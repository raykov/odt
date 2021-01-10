package odt

type Meta struct{}

const metaFile = "meta.xml"

func (m *Meta) Write(w zipWriter) error {
	f, err := w.Create(metaFile)
	if err != nil {
		return err
	}

	_, err = f.Write(m.meta())
	if err != nil {
		return err
	}

	return nil
}

func (m *Meta) meta() []byte {
	return []byte(`<?xml version="1.0" encoding="UTF-8"?>
<office:document-meta 
        xmlns:office="urn:oasis:names:tc:opendocument:xmlns:office:1.0"
        xmlns:xlink="http://www.w3.org/1999/xlink"
        xmlns:dc="http://purl.org/dc/elements/1.1/"
        xmlns:meta="urn:oasis:names:tc:opendocument:xmlns:meta:1.0"
        xmlns:ooo="http://openoffice.org/2004/office"
        xmlns:grddl="http://www.w3.org/2003/g/data-view#" office:version="1.2">
  <office:meta>
    <meta:document-statistic meta:table-count="0" meta:image-count="0" meta:object-count="0" meta:page-count="0" meta:paragraph-count="0" meta:word-count="0" meta:character-count="0" meta:non-whitespace-character-count="0"/>
    <meta:generator>Go ODT generator</meta:generator>
  </office:meta>
</office:document-meta>`)
}
