package odt

type Pictures struct {
	Items []*Image
}

const picturesFilePath = "Pictures/"

func (p *Pictures) Write(w zipWriter) error {
	for _, item := range p.Items {
		path := picturesFilePath + item.FileName

		f, err := w.Create(path)
		if err != nil {
			return err
		}

		_, err = f.Write(item.blob)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *Pictures) Add(images ...*Image) {
	if p.Items == nil {
		p.Items = make([]*Image, 0, len(images))
	}

	p.Items = append(p.Items, images...)
}
