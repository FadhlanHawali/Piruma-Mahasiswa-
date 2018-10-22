package controller

import (
	"github.com/gin-gonic/gin"
	"Piruma/model"
	"net/http"
	"strconv"
	"time"
)

func (idb *InDB) AddHistory (c *gin.Context){
	var(
		add_history model.AddHistory
		history model.History
		result gin.H
	)

	if err:= c.Bind(&add_history); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	history.IdMahasiswa = c.MustGet("id").(string)
	timestamp := strconv.FormatInt(time.Now().Unix(),10)
	history.IdPemesanan = "Ord" +"-"+ string(timestamp)
	history.TimestampStart = add_history.TimestampStart
	history.Telepon = add_history.Telepon
	history.Ruangan = add_history.Ruangan
	history.Keterangan = add_history.Keterangan
	history.Departemen = add_history.Departemen
	history.PenanggungJawab = add_history.PenanggungJawab
	history.StatusPeminjaman = add_history.StatusPeminjaman
	history.StatusSurat = add_history.StatusSurat
	history.TimestampPeminjaman = add_history.TimestampPeminjaman
	history.TimestapEnd = add_history.TimestapEnd

	idb.DB.Create(&history)
	result = gin.H{
		"status":"Success",
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) ListHistory (c *gin.Context){
	var(
		history [] model.History
		list [] model.ListHistory
		result gin.H
	)


	if err := idb.DB.Where("id_mahasiswa = ?",c.MustGet("id").(string)).Find(&history).Error;
		err != nil{
		result = gin.H{
			"result":"Pemesanan tidak ada",
		}
		c.JSON(http.StatusBadRequest,result)
		return
	}else {

		list =make([]model.ListHistory, len(history))
		for i:=0;i<len(history);i++{
			list[i].IdPemesanan = history[i].IdPemesanan
			list[i].TimestapEnd = history[i].TimestapEnd
			list[i].StatusSurat = history[i].StatusSurat
			list[i].StatusPeminjaman = history[i].StatusPeminjaman
			list[i].Departemen = history[i].Departemen
			list[i].Ruangan = history[i].Ruangan
			list[i].TimestampStart = history[i].TimestampStart
		}
		result = gin.H{
			"result":list ,
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