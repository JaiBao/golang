 // routes.go

package routes

import (
	"github.com/labstack/echo/v4"
	"my-echo-project/handlers"
)

func LoadRoutes(e *echo.Echo) {
	e.GET("/", handlers.HelloWorld)
}

