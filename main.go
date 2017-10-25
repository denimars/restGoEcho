package main

import "fmt"
import "restGoEcho/database"

func main(){
	database.DB.Close()
	fmt.Println("test")
}