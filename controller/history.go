package controller

import (
	"github.com/gin-gonic/gin"
	"Piruma/model"
	"net/http"
)

//func (idb *InDB) AddHistory (c *gin.Context){
//	var(
//		add_history model.AddHistory
//		history model.History
//		result gin.H
//	)
//
//	if err:= c.Bind(&add_history); err != nil{
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	history.IdMahasiswa = c.MustGet("id").(string)
//	timestamp := strconv.FormatInt(time.Now().Unix(),10)
//	history.IdPemesanan = "Ord" +"-"+ string(timestamp)
//	history.TimestampStart = add_history.TimestampStart
//	history.Telepon = add_history.Telepon
//	history.Ruangan = add_history.Ruangan
//	history.Keterangan = add_history.Keterangan
//	history.Departemen = add_history.Departemen
//	history.PenanggungJawab = add_history.PenanggungJawab
//	history.TimestapEnd = add_history.TimestapEnd
//
//	idb.DB.Create(&history)
//	result = gin.H{
//		"status":"Success",
//	}
//	c.JSON(http.StatusOK, result)
//}

func (idb *InDB) ListHistory (c *gin.Context){
	var(
		history [] model.History
		result gin.H
	)


	idb.DB.LogMode(true)
	if err := idb.DB.Where("id_mahasiswa = ?",c.MustGet("id").(string)).Find(&history).Error;
		err != nil{
		result = gin.H{
			"result":"Pemesanan tidak ada",
		}
		c.JSON(http.StatusBadRequest,result)
		return
	}else {
		result = gin.H{
			"result":history ,
			"count":  len(history),
		}
		c.JSON(http.StatusOK,result)
	}
}

func (idb *InDB) DetailHistory (c *gin.Context){
	var(
		history model.History
		result gin.H
	)

	idPemesanan := c.Query("idPemesanan")

	if err := idb.DB.Where("id_pemesanan = ?",idPemesanan).First(&history).Error;
	err != nil{
		result = gin.H{
			"result":"Pemesanan tidak ada",
		}
		c.JSON(http.StatusBadRequest,result)
		return
	}else {
		c.JSON(http.StatusOK,history)
	}

}