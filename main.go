package main

import (
	"echoPoc/Server"
	"github.com/labstack/echo/v4"
	"net/http"
)


func main() {
	e := echo.New()
	Server.Routes(e)
	http.Handle("/",e)
	e.Logger.Fatal(e.Start(":4000"))
}
