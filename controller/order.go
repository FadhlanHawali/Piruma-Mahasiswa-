package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"bytes"
	"Piruma/model"
	"encoding/json"
)

func (idb *InDB) AddOrder (c *gin.Context){
	var(
		order model.AddOrder
		timeStamp model.TimeStamp
		statusRuangan model.StatusRuangan
	)

	if err:= c.Bind(&order); err!= nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}

	statusRuangan.StatusPeminjaman = bool(false)
	statusRuangan.StatusSurat = "Surat Belum Dimasukkan"
	timeStamp = order.TimeStamp

	message := map[string]interface{}{
		"id_ruangan":order.IdRuangan,
		"id_departemen":order.IdDepartemen,
		"ruangan":order.Ruangan,
		"departemen":order.Departemen,
		"penanggung_jawab":order.PenanggungJawab,
		"telepon":order.Telepon,
		"keterangan":order.Keterangan,
		"email":order.Email,
		"StatusSurat":map[string]interface{}{
			"status_peminjaman":statusRuangan.StatusPeminjaman,
			"status_surat":statusRuangan.StatusSurat,
		},
		"TimeStamp":map[string]string{
			"timestamp_start":timeStamp.TimestampStart,
			"timestamp_end":timeStamp.TimestampEnd,
		},
	}

	bytesRepresentation, err := json.Marshal(message)
	if err!= nil{
		log.Fatalln(err)
	}

	resp,err := http.Post("https://dteti.au-syd.mybluemix.net/api/addOrder","application/json",bytes.NewBuffer(bytesRepresentation))
	if err!= nil{
		log.Fatalln(err)
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	log.Println(result)
	//log.Println(result["data"])
	c.JSON(http.StatusOK,result)
	return
}

