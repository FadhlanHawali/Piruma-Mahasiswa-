package controller

import (
	"github.com/gin-gonic/gin"
	"sync"
	"fmt"
	"net/http"
	"math/rand"
	"time"
	"Piruma/model"
)
var x  = 0
var bank = [4]model.Bank{
	{
		Name:"Raisul",
		Saldo:100000,
	},
	{
		Name:"Fadhlan",
		Saldo:100000,
	},
	{
		Name:"Ecak",
		Saldo:100000,
	},
	{
		Name:"Penerima",
		Saldo:0,
	},
}


func (idb *InDB) Pembayaran (c * gin.Context){
	var (
		result gin.H
	)
	//idb.DB.LogMode(true)

	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 10; i++ {
		w.Add(1)
		go increment(&w, &m, idb,c)
	}
	w.Wait()
	var total = 0
	fmt.Println("final value of x", x)
	for i:= 0;i<4;i++ {
		fmt.Println(bank[i])
		total = total + bank[i].Saldo
	}

	result = gin.H{
		"final value":x,
		"Bank":bank,
	}
	c.JSON(http.StatusOK, result)
}

func increment(wg *sync.WaitGroup, m *sync.Mutex, idb *InDB, c *gin.Context) {

	m.Lock()
	//time.Sleep(1*time.Second)
	x = x + 1
	myrand := random(0, 3)
	transaksi := random(0,10)
	//fmt.Println("From : " , akun[myrand].Name)
	//fmt.Println("Transaksi : ",transaksi )
	bank[myrand].Saldo -= transaksi
	bank[3].Saldo += transaksi
	if err:=idb.DB.Table("banks").Where("name = ?",bank[myrand].Name).UpdateColumn("saldo", bank[myrand].Saldo).Error;err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}

	if err:=idb.DB.Table("banks").Where("name = ?",bank[3].Name).UpdateColumn("saldo", bank[3].Saldo).Error;err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}
	m.Unlock()
	wg.Done()
}



func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}
