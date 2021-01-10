package odt

func NewTable(rows ...*TableRow) *Table {
	return &Table{
		Rows: rows,
	}
}

type Table struct {
	Rows      []*TableRow
	StyleName string
}

//  <table:table-column table:style-name="Table3.A"/>
//  <table:table-column table:style-name="Table3.B"/>
func (t *Table) String() string {
	content := ""
	for _, element := range t.Rows {
		content += element.String()
	}

	columns := ""
	if len(t.Rows) > 0 {
		for range t.Rows[0].Cells {
			columns += `<table:table-column table:style-name="Default"/>`
		}
	}

	style := "Default"
	if t.StyleName != "" {
		style = t.StyleName
	}

	return `<table:table table:name="" table:style-name="` + style + `">
` + columns + `
<text:soft-page-break/>
` + content + `
</table:table>`
}

func (t *Table) Write() []byte {
	return []byte(t.String())
}

func NewTableRow(cells ...*TableCell) *TableRow {
	return &TableRow{
		Cells: cells,
	}
}

type TableRow struct {
	Cells     []*TableCell
	StyleName string
}

func (r *TableRow) String() string {
	content := ""
	for _, element := range r.Cells {
		content += element.String()
	}

	style := "Default"
	if r.StyleName != "" {
		style = r.StyleName
	}

	return `<table:table-row table:style-name="` + style + `">` + content + `</table:table-row>`
}

func (r *TableRow) WithStyle(styleName string) *TableRow {
	r.StyleName = styleName

	return r
}

func NewTextTableCell(text string) *TableCell {
	return &TableCell{
		Items: []DocElement{NewTextParagraph(text)},
	}
}

type TableCell struct {
	Items     []DocElement
	StyleName string
}

func (c *TableCell) String() string {
	content := ""
	for _, element := range c.Items {
		content += element.String()
	}

	style := "Default"
	if c.StyleName != "" {
		style = c.StyleName
	}

	return `<table:table-cell table:style-name="` + style + `" office:value-type="string">` + content + `</table:table-cell>`
}

func (c *TableCell) WithStyle(styleName string) *TableCell {
	c.StyleName = styleName

	return c
}

/*

<table:table table:name="Table3" table:style-name="Table3">
        <table:table-column table:style-name="Table3.A"/>
        <table:table-column table:style-name="Table3.B"/>
        <text:soft-page-break/>

        <table:table-row table:style-name="Table3.1">
          <table:table-cell table:style-name="Table3.A1" office:value-type="string">
            <text:p text:style-name="P5">
              <text:span text:style-name="T5">Config</text:span>
            </text:p>
          </table:table-cell>
          <table:table-cell table:style-name="Table3.A1" office:value-type="string">
            <text:p text:style-name="P5">
              <text:span text:style-name="T5">Value</text:span>
            </text:p>
          </table:table-cell>
        </table:table-row>

        <table:table-row table:style-name="Table3.1">
          <table:table-cell table:style-name="Table3.A2" office:value-type="string">
            <text:p text:style-name="P5">
              <text:span text:style-name="T5">HPA Min Replicas</text:span>
            </text:p>
          </table:table-cell>
          <table:table-cell table:style-name="Table3.B2" office:value-type="string">
            <text:p text:style-name="P5">
              <text:span text:style-name="T5">4</text:span>
            </text:p>
          </table:table-cell>
        </table:table-row>
        <table:table-row table:style-name="Table3.1">
          <table:table-cell table:style-name="Table3.A3" office:value-type="string">
            <text:p text:style-name="P5">
              <text:span text:style-name="T5">HPA Max Replicas</text:span>
            </text:p>
          </table:table-cell>
          <table:table-cell table:style-name="Table3.B3" office:value-type="string">
            <text:p text:style-name="P5">
              <text:span text:style-name="T5">250 </text:span>
            </text:p>
          </table:table-cell>
        </table:table-row>
        <table:table-row table:style-name="Table3.1">
          <table:table-cell table:style-name="Table3.A4" office:value-type="string">
            <text:p text:style-name="P5">
              <text:span text:style-name="T5">HPA target CPU Utilization Percentage</text:span>
            </text:p>
          </table:table-cell>
          <table:table-cell table:style-name="Table3.B4" office:value-type="string">
            <text:p text:style-name="P5">
              <text:span text:style-name="T5">70%</text:span>
            </text:p>
          </table:table-cell>
        </table:table-row>
        <table:table-row table:style-name="Table3.1">
          <table:table-cell table:style-name="Table3.A5" office:value-type="string">
            <text:p text:style-name="P5">
              <text:span text:style-name="T5">Requested CPU</text:span>
            </text:p>
          </table:table-cell>
          <table:table-cell table:style-name="Table3.B5" office:value-type="string">
            <text:p text:style-name="P5">
              <text:span text:style-name="T5">192m</text:span>
            </text:p>
          </table:table-cell>
        </table:table-row>
        <table:table-row table:style-name="Table3.1">
          <table:table-cell table:style-name="Table3.A6" office:value-type="string">
            <text:p text:style-name="P5">
              <text:span text:style-name="T5">Limit CPU</text:span>
            </text:p>
          </table:table-cell>
          <table:table-cell table:style-name="Table3.B6" office:value-type="string">
            <text:p text:style-name="P5">
              <text:span text:style-name="T5">1</text:span>
            </text:p>
          </table:table-cell>
        </table:table-row>
        <table:table-row table:style-name="Table3.1">
          <table:table-cell table:style-name="Table3.A7" office:value-type="string">
            <text:p text:style-name="P5">
              <text:span text:style-name="T5">Requested Memory</text:span>
            </text:p>
          </table:table-cell>
          <table:table-cell table:style-name="Table3.B7" office:value-type="string">
            <text:p text:style-name="P5">
              <text:span text:style-name="T5">700Mi</text:span>
            </text:p>
          </table:table-cell>
        </table:table-row>
        <table:table-row table:style-name="Table3.1">
          <table:table-cell table:style-name="Table3.A8" office:value-type="string">
            <text:p text:style-name="P5">
              <text:span text:style-name="T5">Limit Memory</text:span>
            </text:p>
          </table:table-cell>
          <table:table-cell table:style-name="Table3.B8" office:value-type="string">
            <text:p text:style-name="P5">
              <text:span text:style-name="T5">700Mi</text:span>
            </text:p>
          </table:table-cell>
        </table:table-row>
        <table:table-row table:style-name="Table3.1">
          <table:table-cell table:style-name="Table3.A9" office:value-type="string">
            <text:p text:style-name="P5">
              <text:span text:style-name="T5">Docker Image</text:span>
            </text:p>
          </table:table-cell>
          <table:table-cell table:style-name="Table3.B9" office:value-type="string">
            <text:p text:style-name="P5">
              <text:span text:style-name="T5">683110685365.dkr.ecr.eu-west-1.amazonaws.com/mobile-api:1.89.74</text:span>
            </text:p>
          </table:table-cell>
        </table:table-row>
      </table:table>

*/
