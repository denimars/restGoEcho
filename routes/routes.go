package routes

import(
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	c "restGoEcho/controllers"
)

func Route(){
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/ptk", c.CreatePtk)
	e.GET("/ptk/:id", c.GetPtk)
	e.PUT("/ptk/:id", c.UpdatePtk)
	e.DELETE("/ptk/:id", c.DeletePtk)



	e.Logger.Fatal(e.Start(":8080"))
}