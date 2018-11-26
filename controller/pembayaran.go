package controller

import (
	"github.com/gin-gonic/gin"
	"sync"
	"fmt"
	"net/http"
	"math/rand"
	"time"
	"Piruma/model"
	"strconv"
	"os"
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


func (idb *InDB) PembayaranLocking (c * gin.Context){
	var (
		result gin.H
	)
	//idb.DB.LogMode(true)

	banyak := c.Query("jumlah")
	jumlah, err := strconv.Atoi(banyak)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

	for i:= 0; i<3;i++{
		if err:=idb.DB.Table("banks").Where("name = ?",bank[i].Name).UpdateColumn("saldo", 100000).Error;err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}
	}
	if err:=idb.DB.Table("banks").Where("name = ?",bank[3].Name).UpdateColumn("saldo", 0).Error;err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}

	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < jumlah; i++ {
		w.Add(1)
		go incrementLocking(&w, &m, idb,c)
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
	return
}

func incrementLocking(wg *sync.WaitGroup, m *sync.Mutex, idb *InDB, c *gin.Context) {

	m.Lock()
	//time.Sleep(1*time.Second)
	x = x + 1
	myrand := random(0, 3)
	//penerima :=
	transaksi := random(0,10)
	fmt.Println("From : " , bank[myrand].Name)
	fmt.Println("Transaksi : ",transaksi )
	bank[myrand].Saldo -= transaksi
	bank[3].Saldo += transaksi
	idb.DB.LogMode(true)
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

func (idb *InDB) PembayaranNonLocking (c * gin.Context){
	var (
		result gin.H
	)
	//idb.DB.LogMode(true)

	banyak := c.Query("jumlah")
	jumlah, err := strconv.Atoi(banyak)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

	for i:= 0; i<3;i++{
		if err:=idb.DB.Table("banks").Where("name = ?",bank[i].Name).UpdateColumn("saldo", 100000).Error;err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}
	}
	if err:=idb.DB.Table("banks").Where("name = ?",bank[3].Name).UpdateColumn("saldo", 0).Error;err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}

	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < jumlah; i++ {
		w.Add(1)
		go incrementNonLocking(&w, &m, idb,c)
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
	return
}

func incrementNonLocking(wg *sync.WaitGroup, m *sync.Mutex, idb *InDB, c *gin.Context) {

	//m.Lock()
	//time.Sleep(1*time.Second)
	x = x + 1
	myrand := random(0, 3)
	//penerima :=
	transaksi := random(0,10)
	fmt.Println("From : " , bank[myrand].Name)
	fmt.Println("Transaksi : ",transaksi )
	bank[myrand].Saldo -= transaksi
	bank[3].Saldo += transaksi
	idb.DB.LogMode(true)
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
	//m.Unlock()
	wg.Done()
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}

