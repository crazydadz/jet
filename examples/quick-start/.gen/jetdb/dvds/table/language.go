//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/postgres"
)

var Language = newLanguageTable()

type languageTable struct {
	postgres.Table

	//Columns
	LanguageID postgres.ColumnInteger
	Name       postgres.ColumnString
	LastUpdate postgres.ColumnTimestamp

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type LanguageTable struct {
	languageTable

	EXCLUDED languageTable
}

// AS creates new LanguageTable with assigned alias
func (a *LanguageTable) AS(alias string) *LanguageTable {
	aliasTable := newLanguageTable()
	aliasTable.Table.AS(alias)
	return aliasTable
}

func newLanguageTable() *LanguageTable {
	return &LanguageTable{
		languageTable: newLanguageTableImpl("dvds", "language"),
		EXCLUDED:      newLanguageTableImpl("", "excluded"),
	}
}

func newLanguageTableImpl(schemaName, tableName string) languageTable {
	var (
		LanguageIDColumn = postgres.IntegerColumn("language_id")
		NameColumn       = postgres.StringColumn("name")
		LastUpdateColumn = postgres.TimestampColumn("last_update")
		allColumns       = postgres.ColumnList{LanguageIDColumn, NameColumn, LastUpdateColumn}
		mutableColumns   = postgres.ColumnList{NameColumn, LastUpdateColumn}
	)

	return languageTable{
		Table: postgres.NewTable(schemaName, tableName, allColumns...),

		//Columns
		LanguageID: LanguageIDColumn,
		Name:       NameColumn,
		LastUpdate: LastUpdateColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
