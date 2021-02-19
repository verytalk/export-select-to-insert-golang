package util

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var dbConn *sql.DB

func getConnection() *sql.DB {
	mysqlUrl := ""+ Config.Mysql.User+":"+ Config.Mysql.Password+"@tcp("+ Config.Mysql.Host+":"+ Config.Mysql.Port+")"+"/"+Config.Mysql.Name
		//connect database
		db, err := sql.Open("mysql", mysqlUrl)
		dbConn = db
		if err != nil {
			log.Fatal(err.Error())
		}

	return dbConn
}



func RunSelect(mSql string) ([]map[string]string,[]string ){
	db := getConnection()
	defer db.Close()

	rows, err := db.Query(mSql)
	if err != nil {
		log.Fatal(err.Error())
	}
	columns, _ := rows.Columns()

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))


	for i := range values {
		scanArgs[i] = &values[i]
	}

	var result []map[string]string
	for rows.Next() {
		res := make(map[string]string)
		_ = rows.Scan(scanArgs...)
		for i, col := range values {
			res[columns[i]] = string(col)
		}
		result = append(result, res)
	}
	_ = rows.Close()

	return result,columns
}



