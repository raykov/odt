package odt

type Settings struct{}

const settingsFile = "settings.xml"

func (s *Settings) Write(w zipWriter) error {
	f, err := w.Create(settingsFile)
	if err != nil {
		return err
	}

	_, err = f.Write(s.settings())
	if err != nil {
		return err
	}

	return nil
}

func (s *Settings) settings() []byte {
	return []byte(`<?xml version="1.0" encoding="UTF-8"?>
<office:document-settings
        xmlns:office="urn:oasis:names:tc:opendocument:xmlns:office:1.0"
        xmlns:xlink="http://www.w3.org/1999/xlink"
        xmlns:config="urn:oasis:names:tc:opendocument:xmlns:config:1.0"
        xmlns:ooo="http://openoffice.org/2004/office" office:version="1.2">
  <office:settings>
    <config:config-item-set config:name="ooo:view-settings">

    </config:config-item-set>
    <config:config-item-set config:name="ooo:configuration-settings">

    </config:config-item-set>
  </office:settings>
</office:document-settings>`)
}
