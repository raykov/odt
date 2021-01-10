package odt

func NewImage(blob []byte, name string) *Image {
	return &Image{
		blob:     blob,
		FileName: name,
	}
}

type Image struct {
	blob      []byte
	FileName  string
	StyleName string
}

func (p *Image) String() string {
	style := "Default"
	if p.StyleName != "" {
		style = p.StyleName
	}

	// svg:height="3.3055in"
	return `<text:p text:style-name="Standard">
<draw:frame draw:style-name="` + style + `" draw:name="` + p.FileName + `" text:anchor-type="as-char" svg:width="6.5in" svg:height="3.3055in" draw:z-index="0">
  <draw:image xlink:href="Pictures/` + p.FileName + `" xlink:type="simple" xlink:show="embed" xlink:actuate="onLoad" loext:mime-type="image/png"/>
</draw:frame>
</text:p>`
}

func (p *Image) Write() []byte {
	return []byte(p.String())
}
