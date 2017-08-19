package sqlite

import (
	"database/sql"
	"fmt"
	"log"
	"bitbucket.org/timabell/sql-data-viewer/schema"
	"strconv"
	"strings"
)

type sqliteModel struct {
	path string
}

func NewSqlite(path string) sqliteModel {
	return sqliteModel{
		path: path,
	}
}

func (model sqliteModel) GetTables() (tables []schema.Table, err error) {
	dbc, err := getConnection(model.path)
	if err != nil {
		return
	}
	defer dbc.Close()

	rows, err := dbc.Query("SELECT name FROM sqlite_master WHERE type='table';")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		rows.Scan(&name)
		tables = append(tables, schema.Table{Schema: "", Name: name})
	}
	return tables, nil
}

func getConnection(path string) (dbc *sql.DB, err error) {
	dbc, err = sql.Open("sqlite3", path)
	if err != nil {
		log.Println("connection error", err)
	}
	return
}

func (model sqliteModel) CheckConnection() (err error) {
	dbc, err := getConnection(model.path)
	if dbc == nil {
		log.Println(err)
		panic("getConnection() returned nil")
	}
	defer dbc.Close()
	tables, err := model.GetTables()
	if err != nil{
		panic(err)
	}
	if len(tables) == 0{
		// https://stackoverflow.com/q/45777113/10245
		panic("No tables found. (Sqlite will create an empty db if the specified file doesn't exist).")
	}
	return
}

func (model sqliteModel) AllFks() (allFks schema.GlobalFkList, err error) {
	tables, err := model.GetTables()
	if err != nil {
		fmt.Println("error getting table list while building global fk list", err)
		return
	}
	allFks = schema.GlobalFkList{}

	// todo: share connection with GetTables()
	dbc, err := getConnection(model.path)
	if err != nil {
		// todo: show in UI
		return
	}
	defer dbc.Close()

	for _, table := range tables {
		allFks[table.String()], err = fks(dbc, table)
		if err != nil {
			// todo: show in UI
			fmt.Println("error getting fks for table "+table.String(), err)
			return
		}
	}
	return
}

func fks(dbc *sql.DB, table schema.Table) (fks schema.FkList, err error) {
	rows, err := dbc.Query("PRAGMA foreign_key_list('" + table.String() + "');")
	if err != nil {
		return
	}
	defer rows.Close()
	fks = schema.FkList{}
	for rows.Next() {
		var id, seq int
		var parentTable, from, to, onUpdate, onDelete, match string
		rows.Scan(&id, &seq, &parentTable, &from, &to, &onUpdate, &onDelete, &match)
		thisRef := schema.Ref{Col: schema.Column{to, ""}, Table: schema.Table{Schema: "", Name: parentTable}}
		fks[schema.Column{from, ""}] = thisRef
	}
	return
}

func (model sqliteModel) GetRows(query schema.RowFilter, table schema.Table, rowLimit int) (rows *sql.Rows, err error) {
	sql := "select * from " + table.String()

	if len(query) > 0 {
		sql = sql + " where "
		clauses := make([]string, 0, len(query))
		for k, v := range query {
			clauses = append(clauses, k+" = "+v[0])
		}
		sql = sql + strings.Join(clauses, " and ")
	}

	if rowLimit > 0 {
		sql = sql + " limit " + strconv.Itoa(rowLimit)
	}

	dbc, err := getConnection(model.path)
	if err != nil {
		log.Println(err)
		panic("GetRows to get connection")
		// todo: show in UI
		return
	}
	defer dbc.Close()

	rows, err = dbc.Query(sql)
	if err != nil {
		log.Println(sql)
		log.Println(err)
		panic("GetRows failed to get query")
		// todo: show in UI
		return
	}
	return
}

func (model sqliteModel) GetColumns(table schema.Table) (cols []schema.Column, err error){
	dbc, err := getConnection(model.path)
	if err != nil {
		log.Println(err)
		panic("GetColumns to get connection")
		// todo: show in UI
		return
	}
	defer dbc.Close()
	rows, err := dbc.Query("PRAGMA table_info('" + table.String() + "');")
	if err != nil {
		return
	}
	defer rows.Close()
	cols = []schema.Column{}
	for rows.Next() {
		var cid int
		var name, typeName string
		var notNull, pk bool
		var defaultValue interface{}
		rows.Scan(&cid, &name, &typeName, &notNull, &defaultValue, &pk)
		thisCol := schema.Column{name, typeName}
		cols = append(cols, thisCol)
	}
	return
}
