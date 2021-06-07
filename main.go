package main

import (
	"shipSlip/router"

	"github.com/labstack/echo/v4"
)

func main() {
	// sqlite, err := services.Connect("./shipSlip.db")
	// if err != nil {
	// 	log.Println(err)
	// }
	// sqlite.CreateTable()
	e := echo.New()
	r := router.New(e)
	r.LineRouting()
	e.Logger.Fatal(e.Start(":1323"))
}
