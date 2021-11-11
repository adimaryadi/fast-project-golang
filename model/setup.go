package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"os"
)


func SetupDB() *gorm.DB {
	USER    := os.Getenv("userMysql")
	PASS    := os.Getenv("passMysql")
	HOST    := os.Getenv("host")
	PORT    := os.Getenv("portMysql")
	DBNAME  := os.Getenv("dbMysql")
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	db, err := gorm.Open("mysql", URL)

	if err != nil {
		panic(err.Error())
	}
	return  db
}