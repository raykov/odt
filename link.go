package odt

const linkDefaultStyleName = "Default"

func NewLink(url string, elements ...DocElement) *Link {
	return &Link{
		URL:  url,
		Text: elements,
	}
}

type Link struct {
	Text             []DocElement
	StyleName        string
	VisitedStyleName string
	URL              string
}

func (l *Link) String() string {
	text := NewSpan(l.URL).String()

	if len(l.Text) > 0 {
		text = ""
		for _, span := range l.Text {
			text += span.String()
		}
	}

	styleName := linkDefaultStyleName
	if l.StyleName != "" {
		styleName = l.StyleName
	}

	visitedStyleName := styleName
	if l.VisitedStyleName != "" {
		visitedStyleName = l.VisitedStyleName
	}

	return `<text:a xlink:type="simple" xlink:href="` +
		l.URL +
		`" text:style-name="` +
		styleName +
		`" text:visited-style-name="` +
		visitedStyleName + `">` +
		text +
		`</text:a>`
}

func (l *Link) Write() []byte {
	return []byte(l.String())
}

func (l *Link) WithStyle(styleNames ...string) *Link {
	if len(styleNames) > 0 {
		l.StyleName = styleNames[0]
	}
	if len(styleNames) > 1 {
		l.VisitedStyleName = styleNames[1]
	}

	return l
}
