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
		//
		//history model.History
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
		"status_proses":false,
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

	//resp,err := http.Post("http://localhost:8080/api/public/addOrder","application/json",bytes.NewBuffer(bytesRepresentation))
	//if err!= nil{
	//	log.Fatalln(err)
	//}
	//
	//resp.Header.Add("Authorization",c.Request.Header.Get("Authorization"))

	req, err := http.NewRequest("POST", "http://localhost:8080/api/public/addOrder", bytes.NewBuffer(bytesRepresentation))
	req.Header.Set("Authorization", c.Request.Header.Get("Authorization"))
	req.Header.Set("Content-Type","application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	//fmt.Println(bytes.NewBuffer(bytesRepresentation))
	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	log.Println(result)
	//log.Println(result["data"])
	c.JSON(http.StatusOK,result)

	//history.Ruangan = order.Ruangan
	//history.StatusSurat = statusRuangan.StatusSurat
	//history.StatusPeminjaman = statusRuangan.StatusPeminjaman
	//history.Keterangan = order.Keterangan
	//history.Telepon = order.Telepon
	//history.PenanggungJawab = order.PenanggungJawab


	return
}

func (idb *InDB) CheckStatus (c *gin.Context){

	idPemesanan := c.Query("idPemesanan")
	//req, _ := http.NewRequest("GET", "localhost:8080/api/order/check?idPemesanan="+idPemesanan, nil)
	//req.Header.Set("Authorization", c.Request.Header.Get("Authorization"))
	//req.Header.Set("Content-Type","application/json")
	//client := &http.Client{}
	//resp, _ := client.Do(req)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "localhost:8080/api/order/check?idPemesanan="+idPemesanan, nil)
	req.Header.Set("Authorization", c.Request.Header.Get("Authorization"))
	resp, _ := client.Do(req)

	//fmt.Println(bytes.NewBuffer(bytesRepresentation))
	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	log.Println(result)
	//log.Println(result["data"])
	c.JSON(http.StatusOK,result)
}

