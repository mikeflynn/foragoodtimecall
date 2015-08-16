package main

import (
	"database/sql"
	"errors"
	"log"
	"reflect"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

type DBRecord interface {
	Save() bool
	Load() bool
}

func Struct2Map(obj DBRecord) map[string]interface{} {
	v := reflect.ValueOf(obj)
	output := make(map[string]interface{}, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Interface() != nil {
			output[v.Type().Field(i).Name] = v.Field(i).Interface()
		}
	}

	return output
}

func MySQLInit() {
	var err error
	DB, err = sql.Open("mysql", "alocaluser:apassword@/foragoodtimecall")
	if err != nil {
		log.Println("SQL Error: " + err.Error())
	}
}

func MySQLClose() {
	DB.Close()
}

func MySQLInsert(table string, values map[string]interface{}) (bool, int64) {
	valStm := make([]string, len(values))
	valCol := make([]string, len(values))
	valArr := make([]interface{}, len(values))

	i := 0
	for c, v := range values {
		valStm[i] = "?"
		valCol[i] = c
		valArr[i] = v

		i++
	}

	query, err := DB.Prepare("INSERT INTO " + table + " (" + strings.Join(valCol, ",") + ") VALUES (" + strings.Join(valStm, ",") + ")")
	if err != nil {
		log.Println("SQL Error: " + err.Error())
		return false, 0
	}

	defer query.Close()

	result, err := query.Exec(valArr...)
	if err != nil {
		log.Println("SQL Error: " + err.Error())
		return false, 0
	}

	newID, err := result.LastInsertId()
	if err != nil {
		log.Println("SQL Error: " + err.Error())
		return true, 0
	}

	return true, newID
}

func MySQLUpdate(table string, values map[string]interface{}, where map[string]interface{}) bool {
	setStr := make([]string, len(values))
	valArr := make([]interface{}, len(values))

	i := 0
	for k, v := range values {
		setStr[i] = k + " = ?"
		valArr[i] = v

		i++
	}

	whereStr, whereVals := formatWhere(where)

	allVals := append(valArr, whereVals...)

	query, err := DB.Prepare("UPDATE " + table + " SET " + strings.Join(setStr, ", ") + " WHERE " + whereStr)

	defer query.Close()

	_, err = query.Exec(allVals...)
	if err != nil {
		log.Println("SQL Error: " + err.Error())
		return false
	}

	return true
}

func MySQLSelect(table string, where map[string]interface{}, fields []string) ([]map[string]string, error) {
	// Build the query
	whereStr, whereVals := formatWhere(where)
	queryStr := "SELECT * FROM " + table + " WHERE " + whereStr

	return MySQLQueryRows(queryStr, whereVals)
}

func MySQLQueryRows(queryStr string, queryVals map[string]interface{}) ([]map[string]string, error) {
	query, err := DB.Prepare(queryStr)

	defer query.Close()

	// Run the query
	rows, err := query.Query(whereVals...)
	if err != nil {
		return nil, errors.New("SQL Select Error 1: " + err.Error())
	}

	// Fetch the query
	columns, err := rows.Columns()
	if err != nil {
		return nil, errors.New("SQL Select Error 2: " + err.Error())
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Process the query results
	results := []map[string]string{}

	for rows.Next() {
		if err := rows.Scan(scanArgs...); err != nil {
			return nil, errors.New("SQL Select Error 3: " + err.Error())
		}

		row := make(map[string]string, len(values))

		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}

			row[columns[i]] = value
		}
		results = append(results, row)
	}

	if err := rows.Err(); err != nil {
		log.Println("SQL Select Error 4: " + err.Error())
		return nil, errors.New("Select failed.")
	}

	return results, nil
}

func formatWhere(where map[string]interface{}) (string, []interface{}) {
	whereStr := ""
	whereVals := []interface{}{}

	for k, v := range where {
		switch v := v.(type) {
		default:
			whereStr = whereStr + k + " = ?"
			whereVals = append(whereVals, v)
		case []interface{}:
			varr := v
			whereStr = whereStr + k + " " + varr[0].(string) + " ?"
			whereVals = append(whereVals, varr[1])
		}
	}

	return whereStr, whereVals
}
