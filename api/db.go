// db.go

package api

import (
    "database/sql"
    "log"

    _ "github.com/go-sql-driver/mysql"
)

// Db 全局數據庫連接對象
var Db *sql.DB

// InitDB 初始化數據庫連接
func InitDB(dataSourceName string) {
    var err error
    Db, err = sql.Open("mysql", dataSourceName)
    if err != nil {
        log.Fatal(err)
    }

    if err := Db.Ping(); err != nil {
        log.Fatal(err)
    }

    log.Println("Connected to the database")
}
