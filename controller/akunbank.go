package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (idb *InDB) AddAkun (c *gin.Context){
	idb.DB.LogMode(true)

	var nama = "Fadhlan"
	var saldo = 4123
	if err:=idb.DB.Table("banks").Where("name = ?",nama).UpdateColumn("saldo", saldo).Error;err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}
}