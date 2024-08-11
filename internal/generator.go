package internal

import "fmt"

type BkgBuilder struct {
	table         *Table
	procedures    []Procedure
	generate_path string
	filePrefix    string
}

func NewBkgBuilder(table *Table) *BkgBuilder {
	return &BkgBuilder{table: table, generate_path: "", procedures: []Procedure{}, filePrefix: "api"}
}

func (g *BkgBuilder) SetFilePrefix(prefix string) *BkgBuilder {
	g.filePrefix = prefix
	return g
}

func (g *BkgBuilder) SetTable(table *Table) *BkgBuilder {
	g.table = table
	// reset procedures
	g.procedures = []Procedure{}
	return g
}

func (g *BkgBuilder) SetGenerateFolder(path string) *BkgBuilder {
	g.generate_path = path
	return g
}

func (g *BkgBuilder) AddGetProcedure() *BkgBuilder {
	proc := NewProcedure(
		"get", g.table,
	).AddParamCols(
		g.table.GetPrimaryCols(),
	).Build()

	g.procedures = append(g.procedures, proc)
	return g
}

func (g *BkgBuilder) AddInsertProcedure() *BkgBuilder {
	proc := NewProcedure(
		"insert", g.table,
	).AddParamCols(
		g.table.GetColumns(),
	).Build()

	g.procedures = append(g.procedures, proc)
	return g
}

func (g *BkgBuilder) AddUpdateProcedure() *BkgBuilder {
	proc := NewProcedure(
		"update", g.table,
	).AddParamCols(
		g.table.GetColumns(),
	).Build()

	g.procedures = append(g.procedures, proc)
	return g
}

func (g *BkgBuilder) Generate() string {
	if g.table == nil {
		panic("table is not set")
	}

	if len(g.procedures) == 0 {
		panic("no procedures added")
	}

	content := g.generatePkgSpic() + "\n\n" + g.generatePkgBody()

	if g.generate_path != "" {
		path := g.generate_path + "/" + g.table.name + "_" + g.filePrefix + ".sql"
		WriteToFile(content, path)
	}

	return content
}

func (g *BkgBuilder) generatePkgSpic() string {
	pkg := fmt.Sprintf("create or replace package %s_api is\n", g.table.name)

	for _, proc := range g.procedures {
		pkg += proc.SpecProcedure(1) + "\n\n"
	}

	pkg += "end;"

	return pkg
}

func (g *BkgBuilder) generatePkgBody() string {
	pkg := fmt.Sprintf("create or replace package body %s_api is\n", g.table.name)

	for _, proc := range g.procedures {
		pkg += proc.BodyProcedure(1) + "\n\n"
	}

	pkg += "end;"

	return pkg
}
