package odt

const paragraphDefaultStyleName = "Default"

func NewParagraph(content ...DocElement) *Paragraph {
	return &Paragraph{
		Content: content,
	}
}

func NewTextParagraph(text string) *Paragraph {
	return &Paragraph{
		Content: []DocElement{NewSpan(text)},
	}
}

func NewLinkParagraph(url, text string) *Paragraph {
	return &Paragraph{
		Content: []DocElement{NewLink(url, NewSpan(text))},
	}
}

type Paragraph struct {
	Content   []DocElement
	StyleName string
}

func (p *Paragraph) Add(elements ...DocElement) *Paragraph {
	if p.Content == nil {
		p.Content = make([]DocElement, 0, len(elements))
	}

	p.Content = append(p.Content, elements...)

	return p
}

func (p *Paragraph) WithStyle(styleName string) *Paragraph {
	p.StyleName = styleName

	return p
}

func (p *Paragraph) String() string {
	content := ""
	for _, elem := range p.Content {
		content += elem.String()
	}

	style := paragraphDefaultStyleName
	if p.StyleName != "" {
		style = p.StyleName
	}

	if content == "" {
		return `<text:p text:style-name="` + style + `"/>`
	}

	return `<text:p text:style-name="` + style + `">` + content + `</text:p>`
}

func (p *Paragraph) Write() []byte {
	return []byte(p.String())
}
