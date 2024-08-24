package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"echo-template/model"
	"echo-template/router/sample"
)

func main() {
	model.Init()

	router := NewRouter()
	router.Logger.Fatal(router.Start(":1232"))
}

func NewRouter() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// e.GET("/sample", sample.GetIndex)
	sample.Router(e.Group("/sample"))

	// e.GET("/sample/users/:id", sample.GetUser)
	// e.GET("/sample/users", sample.GetUserQuery)
	// e.POST("/sample/users", sample.PostUser)

	// e.GET("/address", address.GetAddress)
	// e.GET("/address/ip", address.GetAddressIP)

	// e.POST("/vxlan", vxlan.PostVxlan)
	// e.GET("/vxlan", vxlan.GetIndex)

	return e
}
