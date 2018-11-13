package controller

import (
	"github.com/gin-gonic/gin"
	"Piruma/model"
	"net/http"
	"github.com/gin-gonic/gin/json"
	"log"
	"bytes"
)

//func (idb *InDB) GetRuangan (c *gin.Context){
//
//	//var(
//	//	search model.GetOrder
//	//	result gin.H
//	//)
//	//
//	//if err:= c.Bind(&search); err != nil{
//	//	c.JSON(http.StatusBadRequest,gin.H{
//	//		"error":err.Error(),
//	//	})
//	//	return
//	//}
//
//	message := map[string]interface{}{
//		"angka_1": "10",
//		"angka_2":  "42",
//		"Deskripsi": map[string]string{
//			"desc": "of course!",
//		},
//	}
//
//
//
//	bytesRepresentation, err := json.Marshal(message)
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	resp, err := http.Post("https://kalkulator.au-syd.mybluemix.net/hitung", "application/json", bytes.NewBuffer(bytesRepresentation))
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	var result map[string]interface{}
//
//	json.NewDecoder(resp.Body).Decode(&result)
//
//	log.Println(result)
//	//log.Println(result["data"])
//	c.JSON(http.StatusOK,result)
//
//}
//

func (idb *InDB) GetRuangan (c * gin.Context){

	var(
		search model.SearchRuangan
		timeStamp model.TimeStamp
	)

	if err:=c.Bind(&search);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}
	timeStamp = search.TimeStamp

	message := map[string]interface{}{
		"kapasitas": search.Kapasitas,
		"TimeStamp": map[string]string{
			"timestamp_start": timeStamp.TimestampStart,
			"timestamp_end":timeStamp.TimestampEnd,
		},
	}
	bytesRepresentation, err := json.Marshal(message)
	if err!= nil{
		log.Fatalln(err)
	}

	resp,err := http.Post("https://dteti.au-syd.mybluemix.net/api/search","application/json",bytes.NewBuffer(bytesRepresentation))
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