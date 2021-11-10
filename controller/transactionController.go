package controller

import (
	"crudMysql/model"
	"crudMysql/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/vigneshuvi/GoDateFormat"
	"net/http"
)

func FindTransaction(c *gin.Context) {
	db  := c.MustGet("db").(*gorm.DB)
	var FindTransaction []model.Transaction
	db.Find(&FindTransaction)
	c.JSON(http.StatusOK, gin.H{"code": "00","data": FindTransaction})
}

func CreateTransaction(c *gin.Context) {
	var input model.Transaction
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	fmt.Println(input.Po_number)
	date      :=   tools.GetToday(GoDateFormat.ConvertFormat("yyyy-mm-dd"))
	PO_number :=   tools.GetToday(GoDateFormat.ConvertFormat("yyyymm"))
	data  :=   model.Transaction{Po_number: PO_number,Po_date: date, Po_price_total: input.Po_price_total, Po_cost_total: input.Po_cost_total}
	db    :=   c.MustGet("db").(*gorm.DB)
	db.Create(&data)
	c.JSON(http.StatusOK,gin.H{"code": "00","data": data})
}

func UpdateTransaction(c *gin.Context) {
	db   := c.MustGet("db").(*gorm.DB)
	var  findModel model.Transaction
	if err := db.Where("id = ?",c.Param("id")).First(&findModel).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "row tidak di temukan"})
		return
	}

	var input model.Transaction
	if err :=  c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updateInput  model.Transaction
	updateInput.Po_number      =  input.Po_number
	updateInput.Po_cost_total  =  input.Po_cost_total
	updateInput.Po_price_total =  input.Po_price_total
	db.Model(&findModel).Updates(updateInput)
	c.JSON(http.StatusOK,gin.H{"code": "00","data": findModel})

}

func DeleteTransaction(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var findTransaction model.Transaction
	if err := db.Where("id = ?",c.Param("id")).First(&findTransaction).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Row tidak ditemukan"})
		return
	}

	db.Delete(&findTransaction)
	c.JSON(http.StatusOK,gin.H{"code": "00","data": "deleted"})
}