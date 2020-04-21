package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type personeller struct {
	Name    string
	Surname string
	Age     int
	Salary  int
}

func main() {
	db, _ := gorm.Open("mysql", "root:root@/deneme2?charset=utf8&parseTime=True&loc=Local")

	db.AutoMigrate(&personeller{})
	//db.DropTableIfExists(&personeller{})
	defer db.Close()
}
