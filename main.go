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

// main.go

package main

import (
	"goTest2/api"
	"fmt"
	"github.com/labstack/echo/v4"

)

func main() {
	fmt.Println("Starting the server...")

	// 初始化數據庫連接
	api.InitDB()
    api.SelectDB()

	// 創建 Echo 實例
	e := echo.New()

	// 加載 API 路由
	api.LoadRoutes(e)

	// 啟動服務
	e.Start(":8080")




}



