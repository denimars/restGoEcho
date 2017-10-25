package models

type Ptk struct{
	Id int `gorm:"primary_key auto_increment" json`
	Nptk int `gorm:"not_null" json`
	Name string `gorm:"varchar:100" json`
	Address string `gorm:"varchar:100" json`
	LoanId Loan `gorm:"ForeignKey:PtkId"`
}