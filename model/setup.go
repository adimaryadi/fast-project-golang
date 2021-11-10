package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type ResultSuccess struct {
	 code      string `json:"code"`
	 data      string `json:"data"`
	 message   string `json:"message"`
} 

func SetupDB() *gorm.DB {
	USER    := "root"
	PASS    := ""
	HOST    := "localhost"
	PORT    := "3306"
	DBNAME  := "tester"
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	db, err := gorm.Open("mysql", URL)

	if err != nil {
		panic(err.Error())
	}
	return  db
}