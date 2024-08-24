package sample

import (
	"github.com/labstack/echo/v4"
)

func Router(g *echo.Group) {
	g.GET("/", GetIndex)
}
