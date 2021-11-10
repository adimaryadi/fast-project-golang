package router

import (
	"crudMysql/controller"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db",db)
	})
	public    :=  r.Group("/api")
	public.POST("/authentication", controller.MiddlewareAuth)
	public.POST("/register",controller.RegisterAuth)
	public.GET("/transaction",controller.FindTransaction)
	public.POST("/transaction/create",controller.CreateTransaction)
	public.PATCH("/transaction/:id",controller.UpdateTransaction)
	public.DELETE("/transaction/:id",controller.DeleteTransaction)
	return r
}