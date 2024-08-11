package internal

import (
	"fmt"
	"strings"
)

type Parser struct {
	tokens  []string
	pos     int
	start   int
	tables  []*Table
	current *Table
}

type StateFn func(*Parser) StateFn

func NewParser(ddl string) *Parser {
	tokens := ToknizeDDL(ddl)
	fmt.Println(tokens)
	return &Parser{tokens: tokens, pos: 0, start: 0, tables: []*Table{}}
}

func (p *Parser) Run() []*Table {
	for state := Initial; state != nil; {
		state = state(p)
	}

	if p.current != nil {
		p.addTable()
	}

	return p.tables
}

func (p *Parser) GetTables() []*Table {
	return p.tables
}

func (p *Parser) newTable(name string) {
	p.current = &Table{name: name, cols: []*Column{}}
}

func (p *Parser) addColumn(col *Column) {
	col.table = p.current
	p.current.cols = append(p.current.cols, col)
}

func (p *Parser) SetPrimaryCols(cols []string) {
	for _, col := range cols {
		p.current.SetPrimary(col)
	}
}

func (p *Parser) addTable() {
	p.tables = append(p.tables, p.current)
}

func (p *Parser) Next() string {
	if p.pos < len(p.tokens) {
		p.pos++
		return p.tokens[p.pos-1]
	}
	return ""
}

func (p *Parser) Peek() string {
	if p.pos < len(p.tokens) {
		return p.tokens[p.pos]
	}
	return ""
}

func (p *Parser) Back() {
	p.pos--
}

func (p *Parser) Reset() {
	p.pos = p.start
}

func (p *Parser) Skip() {
	p.start = p.pos
}

func (p *Parser) SkipUntil(token string) {
	for p.Peek() != token {
		p.Next()
	}
	p.Next()
}

func Initial(p *Parser) StateFn {
	for {
		switch p.Next() {
		case "create":
			return CreateTable
		case "":
			return nil
		}
	}
}

func CreateTable(p *Parser) StateFn {
	if p.current != nil {
		p.addTable()
	}
	for {
		switch p.Next() {
		case "table":
			return TableName
		case "or", "replace":
			return CreateTable
		case "":
			return nil
		}
	}
}

func TableName(p *Parser) StateFn {
	name := p.Next()
	if name == "" {
		return nil
	}

	p.newTable(name)

	for {
		switch p.Next() {
		case "(":
			return ColumnState
		case "":
			return nil
		}
	}
}

func ColumnState(p *Parser) StateFn {
	colName := p.Next()
	if colName == "" {
		return nil
	}
	if strings.ToUpper(colName) == "CONSTRAINT" {
		return ConstraintState
	}
	col := &Column{name: colName, isPrimary: false, isUnique: false}

	for {
		switch p.Next() {
		case "(":
			p.SkipUntil(")")
		case "primary":
			col.isPrimary = true
			p.Skip()
		case "unique":
			p.Skip()
			col.isUnique = true
		case ",":
			p.addColumn(col)
			return ColumnState
		case ")":
			p.addColumn(col)
			return CreateTable
		case "":
			return nil
		}
	}
}

func ConstraintState(p *Parser) StateFn {
	primaryCols := make([]string, 0)
	for {
		switch p.Next() {
		case "primary":
			p.SkipUntil("(")
		searchPrimaries:
			for {
				switch p.Next() {
				case ")":
					break searchPrimaries
				case ",":
					continue searchPrimaries
				case "":
					return nil
				default:
					primaryCols = append(primaryCols, p.Peek())
				}
			}
			p.SetPrimaryCols(primaryCols)
		case "(":
			p.SkipUntil(")")
		case ",":
			return ColumnState
		case ")":
			return CreateTable
		case "":
			return nil
		}
	}
}

func ToknizeDDL(ddl string) []string {
	var tokens []string
	var token string
	ddl = strings.NewReplacer(
		"\n", " ",
		"\r", " ",
		"\t", " ",
		"(", " ( ",
		")", " ) ",
		",", " , ",
	).Replace(strings.ToLower(ddl))

	for _, c := range ddl {
		if c == ' ' {
			if token != "" {
				tokens = append(tokens, token)
				token = ""
			}
			continue
		}
		token += string(c)
	}
	if token != "" {
		tokens = append(tokens, token)
	}
	return tokens
}
