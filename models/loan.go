package models

type Loan struct{
	Id int `gorm:"primary_key;auto_incement" json`
	PtkId int `gorm:"not_null" json`
	Value int `json`

}