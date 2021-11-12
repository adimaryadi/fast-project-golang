package migration

import "crudMysql/model"

func Execution()  {
	db := model.SetupDB()
	db.AutoMigrate(&model.Transaction{})
	db.AutoMigrate(&model.Authentification{})
	db.AutoMigrate(&model.SessionToken{})
}