package database

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	mo "restGoEcho/models"
	"runtime"
	"path"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"fmt"
)

type conf struct{
	Database string `yaml:"database"`
	User string `yaml:"user"`
	Password string `yaml:"password"`
}

var DB *gorm.DB

func(c *conf) getConf() *conf{
	_, filename,_,_ := runtime.Caller(0)
	//log.Println(filename)
	filepath := path.Join(path.Dir(filename), "config.yaml")
	//log.Println(filepath)
	yamlfile, err := ioutil.ReadFile(filepath)
	if err != nil{
		log.Printf("yaml file get err %v", err)
	}
	err = yaml.Unmarshal(yamlfile, c)
	if err != nil{
		log.Fatalf("Unmarshall: %v", err)
	}
	return c
}

func init(){
	var err error
	var c conf
	c.getConf()
	//conB := c.User+":"+c.Password+"@/"+c.Database+"?charset=utf8&parseTime=True&loc=Local"
	conB := fmt.Sprintf("%v:%v@/%v?charset=utf8&parseTime=True&loc=Local", c.User, c.Password, c.Database)
	DB, err = gorm.Open("mysql",conB)
	if err != nil{
		log.Println(err)
	}else{
		log.Println("Connected... horraaaay")
	}
	DB.SingularTable(true)
	DB.AutoMigrate(&mo.Loan{}, &mo.Ptk{})
}