package controller

import (
	"github.com/gin-gonic/gin"
	"Piruma/model"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func (idb *InDB) Login(c *gin.Context){
	var (
		login model.Login
		mahasiswa model.Mahasiswa
		result gin.H
	)

	if err:= c.Bind(&login); err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}

	if err:= idb.DB.Where(&model.Mahasiswa{Email:login.Email}).First(&mahasiswa).Error;err!=nil{
		result = gin.H{
			"result":"Email doesn't exist",
			"email":err,
		}
		c.JSON(http.StatusBadRequest,result)
		return
	}else {
		var password_tes = bcrypt.CompareHashAndPassword([]byte(mahasiswa.Password), []byte(login.Password))
		if password_tes == nil {
			sign := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"email":mahasiswa.Email,
				"id_mahasiswa":mahasiswa.IdMahasiswa,
			})
			token, err := sign.SignedString([]byte("secret"))
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": err.Error(),
				})
				c.Abort()
			}
			//login success
			result = gin.H{
				"result":"Success",
				"token":token,
			}
			c.JSON(http.StatusOK,result)
			return
		} else {
			//login failed
			result = gin.H{
				"result":"Salah bos q",
			}
			c.JSON(http.StatusOK,result)
			return
		}
	}
}
