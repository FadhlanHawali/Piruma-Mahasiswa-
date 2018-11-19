package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"Piruma/config"
	"Piruma/controller"
	"Piruma/middleware"
)

func main(){
	db := config.DBInit()
	inDB := &controller.InDB{DB: db}

	router := gin.Default()
	router.POST("/api/signup", inDB.SignUp)
	router.POST("/api/login",inDB.Login)
	router.POST("/api/ruangan/search",inDB.GetRuangan)
	router.POST("/api/ruangan/listroom",inDB.GetListRoom)
	router.POST("api/ruangan/schedule",inDB.GetScheduleRoom)
	router.POST("/api/add_history",middleware.Auth,inDB.AddHistory)
	router.GET("/api/list_history",middleware.Auth,inDB.ListHistory)
	router.GET("/api/history",middleware.Auth,inDB.DetailHistory)
	router.POST("/api/order",inDB.AddOrder)
	router.GET("/api/cobaRandom",inDB.Pembayaran)
	router.GET("/api/tambahAkun",inDB.AddAkun)
	router.Run(":8080")
}
