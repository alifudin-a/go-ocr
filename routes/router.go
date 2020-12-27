package routes

import (
	"github.com/alifudin-a/go-ocr/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Init : initialize router
func Init() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	v1 := e.Group("/api/v1")
	v1.POST("/ocr_translate/id", api.IDtoEN)
	v1.POST("/ocr_translate/en", api.ENtoID)

	return e
}
