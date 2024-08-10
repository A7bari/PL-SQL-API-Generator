package internal

type Column struct {
	name      string
	isPrimary bool
	isUnique  bool
	table     *Table
}

type Table struct {
	name string
	cols []*Column
}

func NewTable(name string) *Table {
	return &Table{name: name, cols: []*Column{}}
}

func (t *Table) AddColumn(name string, isPrimary bool, isUnique bool) {
	t.cols = append(t.cols, &Column{name, isPrimary, isUnique, t})
}

func (t *Table) GetColumns() []*Column {
	var cols []*Column
	cols = append(cols, t.cols...)
	return cols
}

func (t *Table) GetPrimaryCols() []*Column {
	var primaryCols []*Column
	for _, col := range t.cols {
		if col.isPrimary {
			primaryCols = append(primaryCols, col)
		}
	}
	return primaryCols
}

func (t *Table) GetUniqueCols() []*Column {
	var uniqueCols []*Column
	for _, col := range t.cols {
		if col.isUnique {
			uniqueCols = append(uniqueCols, col)
		}
	}
	return uniqueCols
}

func (t *Table) SetPrimary(col string) {
	for _, c := range t.cols {
		if c.name == col {
			c.isPrimary = true
		}
	}
}

func (t *Table) GetRowType() string {
	return t.name + "%rowtype"
}

func (c *Column) GetParamName() string {
	return "p_" + c.name
}

func (c *Column) GetParamType() string {
	return c.table.name + "." + c.name + "%type"
}
