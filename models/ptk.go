package models

type Ptk struct{
	Id int `gorm:"primary_key auto_increment" json`
	Nama string `gorm:"varchar:100" json`
}