package odt

type Content struct {
	Elements    []DocElement
	PreRendered []byte
}

const contentFile = "content.xml"

func (c *Content) Write(w zipWriter) error {
	f, err := w.Create(contentFile)
	if err != nil {
		return err
	}

	_, err = f.Write(c.write())
	if err != nil {
		return err
	}

	return nil
}

func (c *Content) Add(elements ...DocElement) *Content {
	if c.Elements == nil {
		c.Elements = make([]DocElement, 0, len(elements))
	}

	c.Elements = append(c.Elements, elements...)

	return c
}

func (c *Content) write() []byte {
	if len(c.PreRendered) > 0 {
		return c.PreRendered
	}

	content := ""
	for _, element := range c.Elements {
		content += element.String()
	}

	return []byte(`<?xml version="1.0" encoding="UTF-8"?>
<office:document-content ` + documentConfigurations + ` office:version="1.2">
  <office:scripts/>
  <office:automatic-styles>
		<style:style style:name="Default" style:family="table">
		  <style:table-properties style:width="6.5in" fo:margin-left="0in" fo:margin-top="0in" fo:margin-bottom="0in" table:align="left"/>
		</style:style>
		<style:style style:name="Default" style:family="table-column">
		  <style:table-column-properties style:column-width="auto"/>
		</style:style>
		<style:style style:name="Default" style:family="table-row">
		  <style:table-row-properties fo:keep-together="auto"/>
		</style:style>
		<style:style style:name="Default" style:family="table-cell">
		  <style:table-cell-properties style:vertical-align="" fo:padding="0.0694in" fo:border="1pt solid #000000"/>
		</style:style>
    <style:style style:name="Grey" style:family="table-cell">
      <style:table-cell-properties style:vertical-align="" fo:background-color="#efefef" fo:padding="0.0694in" fo:border="1pt solid #000000">
        <style:background-image/>
      </style:table-cell-properties>
    </style:style>
  </office:automatic-styles>
  <office:body>
    <office:text text:use-soft-page-breaks="true">` + content + `</office:text>
  </office:body>
</office:document-content>`)
}

const documentConfigurations = `
        xmlns:office="urn:oasis:names:tc:opendocument:xmlns:office:1.0"
        xmlns:style="urn:oasis:names:tc:opendocument:xmlns:style:1.0"
        xmlns:text="urn:oasis:names:tc:opendocument:xmlns:text:1.0"
        xmlns:table="urn:oasis:names:tc:opendocument:xmlns:table:1.0"
        xmlns:draw="urn:oasis:names:tc:opendocument:xmlns:drawing:1.0"
        xmlns:fo="urn:oasis:names:tc:opendocument:xmlns:xsl-fo-compatible:1.0"
        xmlns:xlink="http://www.w3.org/1999/xlink"
        xmlns:dc="http://purl.org/dc/elements/1.1/"
        xmlns:meta="urn:oasis:names:tc:opendocument:xmlns:meta:1.0"
        xmlns:number="urn:oasis:names:tc:opendocument:xmlns:datastyle:1.0"
        xmlns:svg="urn:oasis:names:tc:opendocument:xmlns:svg-compatible:1.0"
        xmlns:chart="urn:oasis:names:tc:opendocument:xmlns:chart:1.0"
        xmlns:dr3d="urn:oasis:names:tc:opendocument:xmlns:dr3d:1.0"
        xmlns:math="http://www.w3.org/1998/Math/MathML"
        xmlns:form="urn:oasis:names:tc:opendocument:xmlns:form:1.0"
        xmlns:script="urn:oasis:names:tc:opendocument:xmlns:script:1.0"
        xmlns:ooo="http://openoffice.org/2004/office"
        xmlns:ooow="http://openoffice.org/2004/writer"
        xmlns:oooc="http://openoffice.org/2004/calc"
        xmlns:dom="http://www.w3.org/2001/xml-events"
        xmlns:xforms="http://www.w3.org/2002/xforms"
        xmlns:xsd="http://www.w3.org/2001/XMLSchema"
        xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
        xmlns:rpt="http://openoffice.org/2005/report"
        xmlns:of="urn:oasis:names:tc:opendocument:xmlns:of:1.2"
        xmlns:xhtml="http://www.w3.org/1999/xhtml"
        xmlns:grddl="http://www.w3.org/2003/g/data-view#"
        xmlns:officeooo="http://openoffice.org/2009/office"
        xmlns:tableooo="http://openoffice.org/2009/table"
        xmlns:drawooo="http://openoffice.org/2010/draw"
        xmlns:calcext="urn:org:documentfoundation:names:experimental:calc:xmlns:calcext:1.0"
        xmlns:loext="urn:org:documentfoundation:names:experimental:office:xmlns:loext:1.0"
        xmlns:field="urn:openoffice:names:experimental:ooo-ms-interop:xmlns:field:1.0"
        xmlns:formx="urn:openoffice:names:experimental:ooxml-odf-interop:xmlns:form:1.0"
        xmlns:css3t="http://www.w3.org/TR/css3-text/"`

const documentFonts = `
 <office:font-face-decls>
    <style:font-face style:name="OpenSymbol" svg:font-family="OpenSymbol" style:font-charset="x-symbol"/>
    <style:font-face style:name="Arial" svg:font-family="Arial" style:font-family-generic="roman" style:font-pitch="variable"/>
    <style:font-face style:name="Liberation Sans" svg:font-family="&apos;Liberation Sans&apos;" style:font-family-generic="swiss" style:font-pitch="variable"/>
    <style:font-face style:name="Arial1" svg:font-family="Arial" style:font-family-generic="system" style:font-pitch="variable"/>
    <style:font-face style:name="Linux Libertine G" svg:font-family="&apos;Linux Libertine G&apos;" style:font-family-generic="system" style:font-pitch="variable"/>
  </office:font-face-decls>
  <office:automatic-styles>
    <style:style style:name="Standard" style:family="paragraph" style:class="text">
      <style:paragraph-properties fo:line-height="115%"/>
    </style:style>
    <style:style style:name="P1" style:family="paragraph" style:parent-style-name="Standard">
      <style:text-properties fo:font-size="12pt" fo:font-weight="bold" style:font-size-asian="12pt" style:font-weight-asian="bold" style:font-size-complex="12pt"/>
    </style:style>
    <style:style style:name="T1" style:family="text">
      <style:text-properties fo:font-size="14pt" style:font-size-asian="14pt" style:font-size-complex="14pt"/>
    </style:style>
    <style:style style:name="T2" style:family="text">
      <style:text-properties fo:font-size="14pt" fo:font-style="italic" fo:font-weight="bold" style:font-size-asian="14pt" style:font-style-asian="italic" style:font-weight-asian="bold" style:font-size-complex="14pt"/>
    </style:style>
  </office:automatic-styles>
`
