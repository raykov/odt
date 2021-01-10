package odt

type Style struct {
	Name   string
	Family string
	Class  string
	Parent *Style
	Next   *Style
	//Properties map[string]StyleProperty

	ParagraphProperties  map[string]string
	TextProperties       map[string]string
	TableProperties      map[string]string
	TableRowProperties   map[string]string
	ListLevelProperties  map[string]string
	PageLayoutProperties map[string]string
}

func (s *Style) Key() string {
	return s.Name
}

func (s *Style) String() string {
	return ""
}

type DefaultStyle struct {
	Style
}

func (ds *DefaultStyle) Key() string {
	return "default-" + ds.Style.Key()
}

//type StyleProperty struct {
//}

//type OStyles map[string]*Style
//
//func (ss *OStyles) Add(s *Style) {
//	(*ss)[s.Key()] = s
//}
