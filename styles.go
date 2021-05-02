package odt

func NewStyles() *Styles {
	return &Styles{}
}

type Styles []StyleElement

const stylesFile = "styles.xml"

func (ss *Styles) Add(styles ...StyleElement) *Styles {

	*ss = append(*ss, styles...)

	return ss
}

func (ss *Styles) Write(w zipWriter) error {
	f, err := w.Create(stylesFile)
	if err != nil {
		return err
	}

	_, err = f.Write(ss.styles())
	if err != nil {
		return err
	}

	return nil
}

func (ss *Styles) styles() []byte {
	return []byte(`<?xml version="1.0" encoding="UTF-8"?>
<office:document-styles
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
        xmlns:css3t="http://www.w3.org/TR/css3-text/" office:version="1.2">
    <office:font-face-decls>
		<style:font-face style:name="OpenSymbol" svg:font-family="OpenSymbol" style:font-charset="x-symbol"/>
		<style:font-face style:name="Arial" svg:font-family="Arial" style:font-family-generic="roman" style:font-pitch="variable"/>
		<style:font-face style:name="Liberation Sans" svg:font-family="&apos;Liberation Sans&apos;" style:font-family-generic="swiss" style:font-pitch="variable"/>
		<style:font-face style:name="Arial1" svg:font-family="Arial" style:font-family-generic="system" style:font-pitch="variable"/>
		<style:font-face style:name="Linux Libertine G" svg:font-family="&apos;Linux Libertine G&apos;" style:font-family-generic="system" style:font-pitch="variable"/>
    </office:font-face-decls>
    <office:styles>
		<style:default-style style:family="graphic">
		  <style:graphic-properties svg:stroke-color="#3465a4" draw:fill-color="#729fcf" fo:wrap-option="no-wrap" draw:shadow-offset-x="0.1181in" draw:shadow-offset-y="0.1181in" draw:start-line-spacing-horizontal="0.1114in" draw:start-line-spacing-vertical="0.1114in" draw:end-line-spacing-horizontal="0.1114in" draw:end-line-spacing-vertical="0.1114in" style:flow-with-text="false"/>
		  <style:paragraph-properties style:text-autospace="ideograph-alpha" style:line-break="strict" style:font-independent-line-spacing="false">
			<style:tab-stops/>
		  </style:paragraph-properties>
		  <style:text-properties style:use-window-font-color="true" style:font-name="Arial" fo:font-size="11pt" fo:language="en" fo:country="none" style:letter-kerning="false" style:font-name-asian="Arial1" style:font-size-asian="11pt" style:language-asian="zh" style:country-asian="CN" style:font-name-complex="Arial1" style:font-size-complex="11pt" style:language-complex="hi" style:country-complex="IN"/>
		</style:default-style>
		<style:default-style style:family="paragraph">
		  <style:paragraph-properties fo:hyphenation-ladder-count="no-limit" style:text-autospace="ideograph-alpha" style:punctuation-wrap="hanging" style:line-break="strict" style:tab-stop-distance="0.5in" style:writing-mode="page"/>
		  <style:text-properties style:use-window-font-color="true" style:font-name="Arial" fo:font-size="11pt" fo:language="en" fo:country="none" style:letter-kerning="false" style:font-name-asian="Arial1" style:font-size-asian="11pt" style:language-asian="zh" style:country-asian="CN" style:font-name-complex="Arial1" style:font-size-complex="11pt" style:language-complex="hi" style:country-complex="IN" fo:hyphenate="false" fo:hyphenation-remain-char-count="2" fo:hyphenation-push-char-count="2"/>
		</style:default-style>
		<style:default-style style:family="table">
		  <style:table-properties table:border-model="collapsing"/>
		</style:default-style>
		<style:default-style style:family="table-row">
		  <style:table-row-properties fo:keep-together="auto"/>
		</style:default-style>

		<style:style style:name="normal" style:family="paragraph" style:default-outline-level="">
		  <style:paragraph-properties fo:text-align="start" style:justify-single-word="false" fo:orphans="2" fo:widows="2" style:writing-mode="lr-tb"/>
		</style:style>
		<style:style style:name="Heading_20_1" style:display-name="Heading 1" style:family="paragraph" style:parent-style-name="normal" style:next-style-name="Standard" style:default-outline-level="" style:class="text">
		  <style:paragraph-properties fo:margin-top="0.278in" fo:margin-bottom="0.0835in" loext:contextual-spacing="false" fo:line-height="100%" fo:keep-together="always" fo:keep-with-next="always"/>
		  <style:text-properties fo:font-size="20pt" style:font-size-asian="20pt" style:font-size-complex="20pt"/>
		</style:style>
		<style:style style:name="Heading_20_2" style:display-name="Heading 2" style:family="paragraph" style:parent-style-name="normal" style:next-style-name="Standard" style:default-outline-level="" style:class="text">
		  <style:paragraph-properties fo:margin-top="0.25in" fo:margin-bottom="0.0835in" loext:contextual-spacing="false" fo:line-height="100%" fo:keep-together="always" fo:keep-with-next="always"/>
		  <style:text-properties fo:font-size="16pt" fo:font-weight="normal" style:font-size-asian="16pt" style:font-weight-asian="normal" style:font-size-complex="16pt"/>
		</style:style>
		<style:style style:name="Heading_20_3" style:display-name="Heading 3" style:family="paragraph" style:parent-style-name="normal" style:next-style-name="Standard" style:default-outline-level="" style:class="text">
		  <style:paragraph-properties fo:margin-top="0.222in" fo:margin-bottom="0.0555in" loext:contextual-spacing="false" fo:line-height="100%" fo:keep-together="always" fo:keep-with-next="always"/>
		  <style:text-properties fo:color="#434343" fo:font-size="14pt" fo:font-weight="normal" style:font-size-asian="14pt" style:font-weight-asian="normal" style:font-size-complex="14pt"/>
		</style:style>
		<style:style style:name="Heading_20_4" style:display-name="Heading 4" style:family="paragraph" style:parent-style-name="normal" style:next-style-name="Standard" style:default-outline-level="" style:class="text">
		  <style:paragraph-properties fo:margin-top="0.1945in" fo:margin-bottom="0.0555in" loext:contextual-spacing="false" fo:line-height="100%" fo:keep-together="always" fo:keep-with-next="always"/>
		  <style:text-properties fo:color="#666666" fo:font-size="12pt" style:font-size-asian="12pt" style:font-size-complex="12pt"/>
		</style:style>
		<style:style style:name="Heading_20_5" style:display-name="Heading 5" style:family="paragraph" style:parent-style-name="normal" style:next-style-name="Standard" style:default-outline-level="" style:class="text">
		  <style:paragraph-properties fo:margin-top="0.1665in" fo:margin-bottom="0.0555in" loext:contextual-spacing="false" fo:line-height="100%" fo:keep-together="always" fo:keep-with-next="always"/>
		  <style:text-properties fo:color="#666666" fo:font-size="11pt" style:font-size-asian="11pt" style:font-size-complex="11pt"/>
		</style:style>
		<style:style style:name="Heading_20_6" style:display-name="Heading 6" style:family="paragraph" style:parent-style-name="normal" style:next-style-name="Standard" style:default-outline-level="" style:class="text">
		  <style:paragraph-properties fo:margin-top="0.1665in" fo:margin-bottom="0.0555in" loext:contextual-spacing="false" fo:line-height="100%" fo:keep-together="always" fo:keep-with-next="always"/>
		  <style:text-properties fo:color="#666666" fo:font-size="11pt" fo:font-style="italic" style:font-size-asian="11pt" style:font-style-asian="italic" style:font-size-complex="11pt"/>
		</style:style>
		<style:style style:name="Title" style:family="paragraph" style:parent-style-name="normal" style:next-style-name="Standard" style:default-outline-level="" style:class="chapter">
		  <style:paragraph-properties fo:margin-top="0in" fo:margin-bottom="0.0417in" loext:contextual-spacing="false" fo:line-height="100%" fo:keep-together="always" fo:keep-with-next="always"/>
		  <style:text-properties fo:font-size="26pt" style:font-size-asian="26pt" style:font-size-complex="26pt"/>
		</style:style>
		<style:style style:name="Subtitle" style:family="paragraph" style:parent-style-name="normal" style:next-style-name="Standard" style:default-outline-level="" style:class="chapter">
		  <style:paragraph-properties fo:margin-top="0in" fo:margin-bottom="0.222in" loext:contextual-spacing="false" fo:line-height="100%" fo:keep-together="always" fo:keep-with-next="always"/>
		  <style:text-properties fo:color="#666666" style:font-name="Arial" fo:font-family="Arial" style:font-family-generic="roman" style:font-pitch="variable" fo:font-size="15pt" fo:font-style="normal" style:font-name-asian="Arial1" style:font-family-asian="Arial" style:font-family-generic-asian="system" style:font-pitch-asian="variable" style:font-size-asian="15pt" style:font-style-asian="normal" style:font-name-complex="Arial1" style:font-family-complex="Arial" style:font-family-generic-complex="system" style:font-pitch-complex="variable" style:font-size-complex="15pt"/>
		</style:style>

		<style:style style:name="Standard" style:family="paragraph" style:class="text">
			<style:paragraph-properties fo:font-size="14pt" style:font-size-asian="14pt" style:font-size-complex="14pt" fo:line-height="115%"/>
		</style:style>
		<style:style style:name="Bold" style:family="paragraph" style:parent-style-name="Standard">
			<style:text-properties fo:font-weight="bold" style:font-weight-asian="bold" />
		</style:style>
		<style:style style:name="Italic" style:family="paragraph" style:parent-style-name="Standard">
			<style:text-properties fo:font-style="italic" style:font-style-asian="italic" />
		</style:style>
		<style:style style:name="BoldItalic" style:family="paragraph" style:parent-style-name="Standard">
			<style:text-properties fo:font-style="italic" style:font-style-asian="italic" fo:font-weight="bold" style:font-weight-asian="bold" />
		</style:style>
		<style:style style:name="Standard" style:family="text">
			<style:text-properties fo:font-size="14pt" style:font-size-asian="14pt" style:font-size-complex="14pt"/>
		</style:style>
		<style:style style:name="Bold" style:family="text">
			<style:text-properties fo:font-weight="bold" style:font-weight-asian="bold"/>
		</style:style>
		<style:style style:name="Italic" style:family="text">
			<style:text-properties fo:font-style="italic" style:font-style-asian="italic"/>
		</style:style>
		<style:style style:name="BoldItalic" style:family="text">
			<style:text-properties fo:font-style="italic" fo:font-weight="bold" style:font-style-asian="italic" style:font-weight-asian="bold"/>
		</style:style>

		<text:list-style style:name="Default">
		  <text:list-level-style-number text:level="1" text:style-name="ListLabel_20_10" style:num-suffix="." style:num-format="1">
			<style:list-level-properties text:list-level-position-and-space-mode="label-alignment">
			  <style:list-level-label-alignment text:label-followed-by="listtab" fo:text-indent="-0.25in" fo:margin-left="0.5in"/>
			</style:list-level-properties>
		  </text:list-level-style-number>
		  <text:list-level-style-number text:level="2" text:style-name="ListLabel_20_11" style:num-suffix="." style:num-format="a" style:num-letter-sync="true">
			<style:list-level-properties text:list-level-position-and-space-mode="label-alignment">
			  <style:list-level-label-alignment text:label-followed-by="listtab" fo:text-indent="-0.25in" fo:margin-left="1in"/>
			</style:list-level-properties>
		  </text:list-level-style-number>
		  <text:list-level-style-number text:level="3" text:style-name="ListLabel_20_12" style:num-suffix="." style:num-format="i">
			<style:list-level-properties text:list-level-position-and-space-mode="label-alignment" fo:text-align="end">
			  <style:list-level-label-alignment text:label-followed-by="listtab" fo:text-indent="-0.25in" fo:margin-left="1.5in"/>
			</style:list-level-properties>
		  </text:list-level-style-number>
		  <text:list-level-style-number text:level="4" text:style-name="ListLabel_20_13" style:num-suffix="." style:num-format="1">
			<style:list-level-properties text:list-level-position-and-space-mode="label-alignment">
			  <style:list-level-label-alignment text:label-followed-by="listtab" fo:text-indent="-0.25in" fo:margin-left="2in"/>
			</style:list-level-properties>
		  </text:list-level-style-number>
		  <text:list-level-style-number text:level="5" text:style-name="ListLabel_20_14" style:num-suffix="." style:num-format="a" style:num-letter-sync="true">
			<style:list-level-properties text:list-level-position-and-space-mode="label-alignment">
			  <style:list-level-label-alignment text:label-followed-by="listtab" fo:text-indent="-0.25in" fo:margin-left="2.5in"/>
			</style:list-level-properties>
		  </text:list-level-style-number>
		  <text:list-level-style-number text:level="6" text:style-name="ListLabel_20_15" style:num-suffix="." style:num-format="i">
			<style:list-level-properties text:list-level-position-and-space-mode="label-alignment" fo:text-align="end">
			  <style:list-level-label-alignment text:label-followed-by="listtab" fo:text-indent="-0.25in" fo:margin-left="3in"/>
			</style:list-level-properties>
		  </text:list-level-style-number>
		  <text:list-level-style-number text:level="7" text:style-name="ListLabel_20_16" style:num-suffix="." style:num-format="1">
			<style:list-level-properties text:list-level-position-and-space-mode="label-alignment">
			  <style:list-level-label-alignment text:label-followed-by="listtab" fo:text-indent="-0.25in" fo:margin-left="3.5in"/>
			</style:list-level-properties>
		  </text:list-level-style-number>
		  <text:list-level-style-number text:level="8" text:style-name="ListLabel_20_17" style:num-suffix="." style:num-format="a" style:num-letter-sync="true">
			<style:list-level-properties text:list-level-position-and-space-mode="label-alignment">
			  <style:list-level-label-alignment text:label-followed-by="listtab" fo:text-indent="-0.25in" fo:margin-left="4in"/>
			</style:list-level-properties>
		  </text:list-level-style-number>
		  <text:list-level-style-number text:level="9" text:style-name="ListLabel_20_18" style:num-suffix="." style:num-format="i">
			<style:list-level-properties text:list-level-position-and-space-mode="label-alignment" fo:text-align="end">
			  <style:list-level-label-alignment text:label-followed-by="listtab" fo:text-indent="-0.25in" fo:margin-left="4.5in"/>
			</style:list-level-properties>
		  </text:list-level-style-number>
		  <text:list-level-style-number text:level="10" style:num-suffix="." style:num-format="1">
			<style:list-level-properties text:list-level-position-and-space-mode="label-alignment">
			  <style:list-level-label-alignment text:label-followed-by="listtab" text:list-tab-stop-position="2.75in" fo:text-indent="-0.25in" fo:margin-left="2.75in"/>
			</style:list-level-properties>
		  </text:list-level-style-number>
		</text:list-style>




    </office:styles>
</office:document-styles>`)
}
