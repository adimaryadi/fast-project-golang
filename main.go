package main

import (
	"crudMysql/model"
	"crudMysql/router"
)

func main() {
	db := model.SetupDB()
	db.AutoMigrate(&model.Transaction{})
	db.AutoMigrate(&model.Authentification{})
	r := router.SetupRouter(db)
	r.Run()
}
