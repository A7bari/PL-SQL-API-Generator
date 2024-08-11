package main

import (
	"github.com/Ahbari-M/PL-SQL-API-Generator/internal"
)

func main() {

	ddl, _ := internal.ReadFile("./ddl.sql")

	parser := internal.NewParser(ddl)

	tables := parser.Run()
	pkgBuilder := internal.NewBkgBuilder(nil)
	pkgBuilder.SetGenerateFolder("./files")

	for _, table := range tables {
		pkgBuilder.SetTable(table)
		pkgBuilder.AddGetProcedure()
		pkgBuilder.AddInsertProcedure()
		pkgBuilder.AddUpdateProcedure()

		pkgBuilder.Generate()

		// fmt.Println("/")
		// fmt.Println()
	}

}
