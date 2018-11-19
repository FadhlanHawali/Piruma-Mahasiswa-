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

func (idb *InDB) GetListRoom (c *gin.Context){

	var (
		searchList model.ListRoom
	)

	if err:= c.Bind(&searchList);err!= nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}
	resp, err:=http.Get("https://dteti.au-syd.mybluemix.net/api/search/"+searchList.IdDepartemen+"?kapasitas="+searchList.Kapasitas)
	if err != nil{
		log.Fatalln(err)
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)
	c.JSON(http.StatusOK,result)
	return
}

func (idb *InDB) GetScheduleRoom (c *gin.Context){
	var(
		scheduleRoom model.ScheduleRoom
		timeStamp model.TimeStamp
	)

	if err:= c.Bind(&scheduleRoom); err!= nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}

	timeStamp = scheduleRoom.TimeStamp

	resp, err:= http.Get("https://dteti.au-syd.mybluemix.net/api/ruangan/"+scheduleRoom.IdRuangan+"/time?start="+timeStamp.TimestampStart+"+&end="+timeStamp.TimestampEnd)
	if err!= nil{
		log.Fatalln(err)
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)
	c.JSON(http.StatusOK,result)
	return
}


