// routes.go

package api

import (
	"github.com/labstack/echo/v4"
)

// LoadRoutes 用於加載 API 路由
func LoadRoutes(e *echo.Echo) {
	e.POST("/add-grade", AddGrade)
	e.GET("/get-grades", GetGrades)
}
