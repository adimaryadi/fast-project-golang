package model

type Authentification struct {
	ID     	   uint  `json:"id" gorm:"primary_key"`
	Username   string `form:"username" json:"username" binding:"required"`
	Password   string `form:"password" json:"password" binding:"required"`
}