package main

import (
	"crudMysql/migration"
	"crudMysql/model"
	"crudMysql/router"
	"crudMysql/tools"
	"os"
)

func main() {
	tools.SetConfig()
	db := model.SetupDB()
	migration.Execution()
	r := router.SetupRouter(db)
	r.Run(":"+os.Getenv("port"))
}
