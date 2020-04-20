package main

import (
	"fmt"

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
	db, err := gorm.Open("mysql", "root:root@/yenideneme?charset=utf8&parseTime=True&loc=Local")
	fmt.Print(err)
	db.AutoMigrate(&personeller{})
	//db.DropTableIfExists(&personeller{})
	defer db.Close()
}
