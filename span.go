package odt

const spanDefaultStyleName = "Default"

func NewSpan(text string, styleName ...string) *Span {
	s := Span{
		Text: text,
	}
	if len(styleName) > 0 {
		s.StyleName = styleName[0]
	}

	return &s
}

type Span struct {
	StyleName string
	Text      string
}

func (s *Span) String() string {
	style := spanDefaultStyleName
	if s.StyleName != "" {
		style = s.StyleName
	}

	return `<text:span text:style-name="` + style + `">` + s.Text + `</text:span>`
}

func (s *Span) Write() []byte {
	return []byte(s.String())
}

func (s *Span) WithStyle(styleName string) *Span {
	s.StyleName = styleName

	return s
}
