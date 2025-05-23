package config

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	// Connect to the database
	d, err := gorm.Open("mysql", "myuser:mypassword@tcp(localhost:3306)/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db=d
}
 
func GetDB() *gorm.DB {
	return db
}
