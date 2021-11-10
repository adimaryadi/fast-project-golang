package controller

import (
	"crudMysql/model"
	"crudMysql/tools"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func FindUsers(c *gin.Context) {
	db  	:= c.MustGet("db").(*gorm.DB)
	var authfind []model.Authentification
	db.Find(&authfind)
}

func MiddlewareAuth(c *gin.Context) {
	var input model.Authentification

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "99", "message": err.Error()})
		return
	}
	db     := c.MustGet("db").(*gorm.DB)
	hashPassword :=  tools.EncryptionSha256([]byte(input.Password))
	if err := db.Where("username = ? AND password = ? ", input.Username,hashPassword).First(&input).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "78","message": "username password salah"})
		return
	}
	token, err := tools.GenerateToken(input.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "08","message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "00","message": token})
}

func RegisterAuth(c * gin.Context) {
	var input model.Authentification

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "99", "message": err.Error()})
		return
	}
	hashPassword :=  tools.EncryptionSha256([]byte(input.Password))
	save 	 	 :=  model.Authentification{Username: input.Username, Password: hashPassword}
	db 	 	 	 :=  c.MustGet("db").(*gorm.DB)
	db.Create(&save)
	c.JSON(http.StatusOK, gin.H{"code": "00","message": "Create users "+input.Username})
}