package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Server(port string) {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world")
	})

	e.Logger.Fatal(e.Start(":" + port))
}