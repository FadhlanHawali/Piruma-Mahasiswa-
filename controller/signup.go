package controller

import (
	"github.com/gin-gonic/gin"
	"Piruma/model"
	"net/http"
	"time"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"fmt"
)

func (idb *InDB) SignUp(c * gin.Context){
	var(
		mahasiswa model.Mahasiswa
		signup model.SignUp
		result gin.H
	)

	if err:= c.Bind(&signup); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := idb.DB.Where(&model.Mahasiswa{Email:signup.Email}).First(&mahasiswa).Error; err!= nil{

	} else {
		result = gin.H{
			"status" : "failed",
			"reason":"Email does exist",
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}


	timestamp := strconv.FormatInt(time.Now().Unix(),10)
	fmt.Println(timestamp)
	mahasiswa.IdMahasiswa = "Id" +"-"+ string(timestamp)
	mahasiswa.Username = signup.Username
	mahasiswa.Email = signup.Email
	mahasiswa.NIM = signup.NIM
	mahasiswa.Telepon = signup.Telepon

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signup.Password), bcrypt.DefaultCost);if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
	}

	mahasiswa.Password = string(hashedPassword)
	idb.DB.Create(&mahasiswa)
	result = gin.H{
		"status":"success",
	}
	c.JSON(http.StatusOK, result)
}