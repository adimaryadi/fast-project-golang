package main

import (
	"crudMysql/model"
	"crudMysql/router"
	"crudMysql/tools"
)

func main() {
	tools.SetConfig()
	db := model.SetupDB()
	db.AutoMigrate(&model.Transaction{})
	db.AutoMigrate(&model.Authentification{})
	r := router.SetupRouter(db)
	r.Run()
}
