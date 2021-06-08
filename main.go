package main

import (
	controllers "shipSlip/controllers"
	"shipSlip/router"

	"github.com/labstack/echo/v4"
)

func init() {
	controllers.ConnectToSqlite()
}

func main() {
	e := echo.New()
	r := router.New(e)
	r.LineRouting()
	e.Logger.Fatal(e.Start(":1323"))
}
