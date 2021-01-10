package odt

import (
	"archive/zip"
	"bytes"
	"os"
)

func NewDocument(styles *Styles) *Document {
	return &Document{
		Styles: styles,

		Components: []DocComponent{
			&Manifest{},
			&Meta{},
			&Mimetype{},
		},
	}
}

type DocComponent interface{ Write(zipWriter) error }

type Document struct {
	Components []DocComponent

	MetaInf  MetaInf
	Pictures Pictures
	Content  Content
	//Manifest Manifest
	//Meta     Meta
	//Mimetype Mimetype
	Settings Settings
	Styles   *Styles
	Fonts    []Font
}

func (d *Document) Bytes() ([]byte, error) {
	buf := new(bytes.Buffer)

	w := zip.NewWriter(buf)

	err := d.write(w)
	if err != nil {
		w.Close()

		return nil, err
	}
	err = w.Close()
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (d *Document) Write(fileName string) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	bs, err := d.Bytes()

	f.Write(bs)

	return nil
}

func (d *Document) Add(elements ...DocElement) *Document {
	d.Content.Add(elements...)

	return d
}

func (d *Document) write(w zipWriter) error {
	for _, c := range d.Components {
		err := c.Write(w)
		if err != nil {
			return err
		}
	}

	err := d.MetaInf.Write(w)
	if err != nil {
		return err
	}
	err = d.Pictures.Write(w)
	if err != nil {
		return err
	}
	err = d.Content.Write(w)
	if err != nil {
		return err
	}
	//err = d.Manifest.Write(w)
	//if err != nil {
	//	return err
	//}
	//err = d.Meta.Write(w)
	//if err != nil {
	//	return err
	//}
	//err = d.Mimetype.Write(w)
	//if err != nil {
	//	return err
	//}
	err = d.Settings.Write(w)
	if err != nil {
		return err
	}
	err = d.Styles.Write(w)
	if err != nil {
		return err
	}

	return nil
}
