package ptk

import(
	"net/http"
	"github.com/labstack/echo"
	db "restGoEcho/database"
	m "restGoEcho/models"
	"log"
	"strconv"
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

func GetPtk(c echo.Context) error{
	var ptk m.Ptk
	id, _ := strconv.Atoi(c.Param("id"))
	sts := db.DB.Find(&ptk, id)
	if sts.RecordNotFound(){
		message := map[string] string{"message":"Not Found"}
		return c.JSON(http.StatusOK, message)
	}
	return c.JSON(http.StatusOK, ptk)
}

func UpdatePtk(c echo.Context) error{
	id, _ := strconv.Atoi(c.Param("id"))
	var ptk m.Ptk
	var ptks m.Ptk
	sts := db.DB.Find(&ptk, id)
	if sts.RecordNotFound(){
		message := map[string] string{"message":"Not Found"}
		return c.JSON(http.StatusOK, message)
	}
	u := &ptks
	if err := c.Bind(u); err != nil{
		return err
	}
	ptk.Name = ptks.Name
	ptk.Address = ptks.Address
	db.DB.Save(&ptk)
	message := map[string] string{"message":"success"}
	return c.JSON(http.StatusOK, message)
}

func DeletePtk(c echo.Context) error{
	id, _ := strconv.Atoi(c.Param("id"))
	var ptk m.Ptk
	err := db.DB.Where("id=?", id).Delete(&ptk)
	if err.Error!=nil{
		log.Println(err.Error)
	}
	//return c.NoContent(http.StatusNoContent)
	message := map[string] string{"status": "success"}
	return c.JSON(http.StatusOK, message)
}
