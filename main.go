package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type staff struct {
	npm   string
	nama  string
	kelas string
	hoby  string
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/projek23")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func sqlQuery() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	//var age = 20
	rows, err := db.Query("select * from tbprojek23")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []staff

	for rows.Next() {
		var each = staff{}
		var err = rows.Scan(&each.npm, &each.nama, &each.kelas, &each.hoby)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		result = append(result, each)
	}

	for _, each := range result {
		fmt.Println(each)
		//fmt.Println(each.matkul)
		//fmt.Println(each.id)
	}
}
func main() {
	sqlQuery()
}
