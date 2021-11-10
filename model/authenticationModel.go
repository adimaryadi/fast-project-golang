package model

type Authentification struct {
	ID     	   uint  `json:"id" gorm:"primary_key"`
	Username   string `form:"username" gorm: 40; unique json:"username" binding:"required"`
	Password   string `form:"password" json:"password" binding:"required"`
}