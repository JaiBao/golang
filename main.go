// package main

// import (
//     "GOTEST2/api"
//     "GOTEST2/config"
//     "github.com/labstack/echo/v4"
//     "github.com/labstack/echo/v4/middleware"
// )

// func main() {
//     e := echo.New()

//     // 中間件，用於處理跨域請求和日誌等
//     e.Use(middleware.CORS())
//     e.Use(middleware.Logger())

//     // 初始化數據庫連接
//     db := config.InitDB()
//     defer db.Close()

//     // 設置路由
//     api.SetupRoutes(e, db)

//     // 啟動服務
//     e.Start(":8080")
// }

// main.go



// main.go

package main

import (
    "github.com/labstack/echo/v4"
    "goTest2/api"
)

func main() {
    // 初始化數據庫連接，替換為您的 MySQL 連接字符串
    api.InitDB("username:password@tcp(127.0.0.1:3306)/grades")

    e := echo.New()

    // 設置路由
    e.POST("/add-grade", api.AddGrade)
    e.GET("/get-grades", api.GetGrades)

    // 啟動服務
    e.Start(":8080")
}


