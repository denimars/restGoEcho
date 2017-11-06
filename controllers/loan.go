package controllers

import(
	"net/http"
	"github.com/labstack/echo"
	db "restGoEcho/database"
	m "restGoEcho/models"
	//"log"
	"strconv"
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

func GetLoan(c echo.Context) error{
	var loan m.Loan
	id, _ := strconv.Atoi(c.Param("id"))
	sts := db.DB.Find(&loan, id)
	if sts.RecordNotFound(){
		message := map[string] string{"message":"Not Found :)"}
		return c.JSON(http.StatusOK, message)
	}
	return c.JSON(http.StatusOK, loan)
}

func UpdateLoan(c echo.Context) error{
	id, _ := strconv.Atoi(c.Param("id"))
	var loan m.Loan
	var loans m.Loan
	sts := db.DB.Find(&loan, id)
	if sts.RecordNotFound(){
		message := map[string] string{"message":"Not Found :)"}
		return c.JSON(http.StatusOK, message)
	}
	u := &loans
	if err := c.Bind(u); err != nil{
		return err
	}
	loan.PtkId = loans.PtkId
	loan.Value = loans.Value
	db.DB.Save(&loan)
	message := map[string] string{"message":"success"}
	return c.JSON(http.StatusOK, message)
}