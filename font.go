package odt

type Font struct {
	Name          string
	Family        string
	FamilyGeneric string
	Charset       string
	Pitch         string
}

func (f *Font) String() string {
	attributes := ""
	if f.FamilyGeneric != "" {
		attributes += ` style:font-family-generic="` + f.FamilyGeneric + `"`
	}

	if f.Charset != "" {
		attributes += ` style:font-charset="` + f.Charset + `"`
	}

	if f.Pitch != "" {
		attributes += ` style:font-pitch="` + f.Pitch + `"`
	}

	return `<style:font-face style:name="` + f.Name + `" svg:font-family="` + f.Family + `" ` + attributes + ` />`
}
