// routes.go

package api

import (
    "github.com/labstack/echo/v4"
    "api/handlers"
)

func LoadRoutes(e *echo.Echo) {
    e.POST("/add-grade", handlers.AddGrade)
    e.GET("/get-grades", handlers.GetGrades)
}
