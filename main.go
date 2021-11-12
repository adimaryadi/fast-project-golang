package main

import (
	"crudMysql/model"
	"crudMysql/router"
	"crudMysql/tools"
	"os"
)

func main() {
	tools.SetConfig()
	db := model.SetupDB()
	db.AutoMigrate(&model.Transaction{})
	db.AutoMigrate(&model.Authentification{})
	db.AutoMigrate(&model.SessionToken{})
	r := router.SetupRouter(db)
	r.Run(":"+os.Getenv("port"))
}
