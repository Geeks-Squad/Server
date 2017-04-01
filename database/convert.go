package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func makeStructJSON(rows *sql.Rows, w *http.ResponseWriter) error {

	defer rows.Close()
	fmt.Println(rows)
	columns, err := rows.Columns()
	if err != nil {
		print(err)
	}

	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)

	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		print(err)
	}

	(*w).Header().Set("Content-Type", "application/json")

	_, err = (*w).Write(jsonData)
	fmt.Println(string(http.StatusInternalServerError))
	if err != nil {
		(*w).Header().Set("Status-Code", string(http.StatusInternalServerError))
		print(err)
	}
	return nil
}
