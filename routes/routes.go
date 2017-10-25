package routes

import(
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Route(){
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())


	e.Logger.Fatal(e.Start(":8080"))
}