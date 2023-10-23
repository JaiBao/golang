// db.go

package api

import (
    "database/sql"
    "fmt"
    "log"
    _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Order 结构体用于表示数据库表中的记录
type Order struct {
    id    int    // 假设数据库表中有一个名为 "ID" 的整数字段

}

func InitDB() {
    dsn := "root:root@tcp(127.0.0.1:8889)/test"
    var err error
    db, err = sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("已連結至資料庫")
}

func SelectDB() ([]Order, error) {
    var orders []Order
    query := "SELECT id FROM student"

    rows, err := db.Query(query)
    if err != nil {
        log.Printf("查询数据库出错：%v", err)
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var order Order
        if err := rows.Scan(&order.id); err != nil {
            log.Printf("扫描数据库行出错：%v", err)
            return nil, err
        }
        orders = append(orders, order)
    }

    log.Println("查询数据库成功")
    return orders, nil
}

