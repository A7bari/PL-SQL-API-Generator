package internal

import (
	"fmt"
	"strings"
)

type Procedure interface {
	SpecProcedure(tabs int) string
	BodyProcedure(tabs int) string
}

type ProcedureBuilder struct {
	tabs      int
	name      string
	table     *Table
	paramCols []*Column
	outRow    bool
}

func NewProcedure(name string, table *Table) *ProcedureBuilder {
	return &ProcedureBuilder{name: name, table: table, tabs: 1, outRow: true}
}

func (p *ProcedureBuilder) SetTabs(tabs int) *ProcedureBuilder {
	p.tabs = tabs
	return p
}

// AddParamCols adds the columns to the default columns of the procedure
// get: filter by columns in addition to the primary columns
// insert: insert the columns
// update: columns to be updated in addition to the primary columns
func (p *ProcedureBuilder) AddParamCols(cols []*Column) *ProcedureBuilder {
	if p.name == "get" {
		// insure the primary columns are all exist first
		p.paramCols = p.table.GetPrimaryCols()
		for _, col := range cols {
			if !col.isPrimary {
				p.paramCols = append(p.paramCols, col)
			}
		}
		return p
	}

	if p.name == "update" {
		// insure the primary columns are all exist first
		p.paramCols = p.table.GetPrimaryCols()
		for _, col := range cols {
			if !col.isPrimary {
				p.paramCols = append(p.paramCols, col)
			}
		}
		return p
	}

	p.paramCols = cols
	return p
}

func (p *ProcedureBuilder) OutRow(out bool) *ProcedureBuilder {
	p.outRow = out
	return p
}

func (p *ProcedureBuilder) getSignature() string {
	tabs := strings.Repeat("\t", p.tabs)
	proc := tabs + fmt.Sprintf("procedure %s (\n", p.name)

	for i, col := range p.paramCols {
		proc += tabs
		if i == 0 {
			proc += fmt.Sprintf("\t%s in %s \n", col.GetParamName(), col.GetParamType())
		} else {
			proc += fmt.Sprintf("\t,%s in %s \n", col.GetParamName(), col.GetParamType())
		}
	}

	if p.outRow {
		proc += tabs + fmt.Sprintf("\t,p_row out %s\n", p.table.GetRowType())
	}

	proc += tabs + ")"
	return proc
}

func (p *ProcedureBuilder) Build() Procedure {
	if p.paramCols == nil {
		panic("procedure: " + p.name + " must have columns")
	}

	if p.name == "insert" && len(p.paramCols) == 0 {
		panic("insert procedure must have columns to insert")
	}

	if p.name == "get" && !p.outRow {
		panic("get procedure must return a row, set outRow to true")
	}

	return p
}

func (p *ProcedureBuilder) SpecProcedure(tabs int) string {
	return p.SetTabs(tabs).getSignature() + ";"
}

func (p *ProcedureBuilder) BodyProcedure(tabs int) string {
	p.SetTabs(tabs)
	switch p.name {
	case "get":
		return getProcedure(p.tabs, p.getSignature(), p.table, p.paramCols, p.outRow)
	case "insert":
		return insertProcedure(p.tabs, p.getSignature(), p.table, p.paramCols, p.outRow)
	case "update":
		return updateProcedure(p.tabs, p.getSignature(), p.table, p.paramCols, p.outRow)
	default:
		return ""
	}
}

func getProcedure(tabs_num int, signature string, table *Table, paramCols []*Column, outRow bool) string {
	tabs := strings.Repeat("\t", tabs_num)

	// create the procedure
	proc := signature + " is\n"

	// declare the row
	if outRow {
		proc += tabs + fmt.Sprintf("\tlrow %s;\n", table.GetRowType())
	}

	// begin the procedure
	proc += tabs + "begin\n"

	// select the row
	proc += tabs + fmt.Sprintf("\tselect * into lrow from %s where ", table.name)
	for i, col := range paramCols {
		if i == 0 {
			proc += fmt.Sprintf("%s = %s", col.name, col.GetParamName())
		} else {
			proc += fmt.Sprintf(" and %s = %s", col.name, col.GetParamName())
		}
	}
	proc += ";\n"

	// return the row
	if outRow {
		proc += tabs + "\tp_row := lrow;\n"
	}

	// end the procedure
	proc += tabs + "end get;"
	return proc
}

func insertProcedure(tabs_num int, signature string, table *Table, paramCols []*Column, outRow bool) string {
	tabs := strings.Repeat("\t", tabs_num)

	// create the procedure
	proc := signature + " is\n"

	// declare the row
	if outRow {
		proc += tabs + fmt.Sprintf("\tlrow %s;\n", table.GetRowType())
	}

	// begin the procedure
	proc += tabs + "begin\n"

	// insert the row
	proc += tabs + fmt.Sprintf("\tinsert into %s (", table.name)
	for i, col := range paramCols {
		if i == 0 {
			proc += col.name
		} else {
			proc += fmt.Sprintf(", %s", col.name)
		}
	}
	proc += ") values ("
	for i, col := range paramCols {
		if i == 0 {
			proc += col.GetParamName()
		} else {
			proc += fmt.Sprintf(", %s", col.GetParamName())
		}
	}
	proc += ")"

	// return the row
	if outRow {
		proc += "\n" + tabs + "\treturning * into lrow;\n"
		proc += tabs + "\tp_row := lrow;\n"
	} else {
		proc += ";\n"
	}

	// end the procedure
	proc += tabs + "end insert;"
	return proc
}

func updateProcedure(tabs_num int, signature string, table *Table, paramCols []*Column, outRow bool) string {
	tabs := strings.Repeat("\t", tabs_num)

	// create the procedure
	proc := signature + " is\n"

	// declare the row
	if outRow {
		proc += tabs + fmt.Sprintf("\tlrow %s;\n", table.GetRowType())
	}

	// begin the procedure
	proc += tabs + "begin\n"

	// update the row
	proc += tabs + fmt.Sprintf("\tupdate %s set ", table.name)
	index := 0
	for _, col := range paramCols {
		if !col.isPrimary {
			if index == 0 {
				proc += fmt.Sprintf("%s = %s", col.name, col.GetParamName())
			} else {
				proc += fmt.Sprintf(", %s = %s", col.name, col.GetParamName())
			}
			index++
		}
	}
	proc += " where "
	for i, col := range table.GetPrimaryCols() {
		if i == 0 {
			proc += fmt.Sprintf("%s = %s", col.name, col.GetParamName())
		} else {
			proc += fmt.Sprintf(" and %s = %s", col.name, col.GetParamName())
		}
	}

	// return the row
	if outRow {
		proc += tabs + "returning * into lrow;\n"
		proc += tabs + "\tp_row := lrow;\n"
	} else {
		proc += ";\n"
	}

	// end the procedure
	proc += tabs + "end update;"
	return proc
}
