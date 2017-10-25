package ptk

import(
	"net/http"
	"github.com/labstack/echo"
	db "restGoEcho/database"
	m "restGoEcho/models"
	//"log"
	//"strconv"
)

func CreatePtk(c echo.Context) error{
	var ptk m.Ptk
	var ptks m.Ptk
	u := &ptks
	if err := c.Bind(u); err != nil{
		return err
	}
	ptk.Id = 0
	ptk.Nptk = u.Nptk
	ptk.Name = u.Name
	ptk.Address = u.Address
	db.DB.Create(&ptk)
	return c.JSON(http.StatusCreated, u)
}