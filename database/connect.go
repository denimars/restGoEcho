package database

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var DB *gorm.DB

func init(){
	var err error
	DB, err = gorm.Open("mysql","root:root@/dev?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		log.Println(err)
	}else{
		log.Println("Connected... horraaaay")
	}
	DB.SingularTable(true)
	//DB.AutoMigrate()
}