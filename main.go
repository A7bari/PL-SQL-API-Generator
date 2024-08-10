package main

import (
	"fmt"

	"github.com/Ahbari-M/PL-SQL-API-Generator/internal"
)

func main() {

	ddl := "CREATE TABLE test (id INT PRIMARY KEY, name VARCHAR(255), age INT); CREATE TABLE test2 (id INT PRIMARY KEY, name VARCHAR(255), age INT);"
	parser := internal.NewParser(ddl)
	parser.Run()

	tables := parser.GetTables()
	pkgBuilder := internal.NewBkgBuilder(nil)
	for _, table := range tables {
		pkgBuilder.SetTable(table)
		pkgBuilder.AddGetProcedure()
		pkgBuilder.AddInsertProcedure()
		pkgBuilder.AddUpdateProcedure()

		fmt.Println(pkgBuilder.Generate())

		fmt.Println("/")
		fmt.Println()
	}

}
