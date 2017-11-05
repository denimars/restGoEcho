package controllers

import(
	"net/http"
	"github.com/labstack/echo"
	db "restGoEcho/database"
	m "restGoEcho/models"
	/**"log"
	"strconv"*/
)

func CreateLoan(c echo.Context) error{
	var loan m.Loan
	var loans m.Loan
	u := &loans
	if err := c.Bind(u); err != nil{
		return err
	}
	loan.Id = 0
	loan.PtkId= u.PtkId
	loan.Value= u.Value
	db.DB.Create(&loan)
	return c.JSON(http.StatusCreated,u)
}