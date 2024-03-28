package controller

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shplume/zhulong/service"
	"github.com/shplume/zhulong/utils"
)

func CreateUser(c *gin.Context) {
	var data struct {
		Account  string `json:"account" binding:"required"`
		Password string `json:"password" binding:"required"`
		Role     int    `json:"role" binding:"required"`
	}

	if err := c.BindJSON(&data); err != nil {
		slog.Error(err.Error())

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Parameter",
		})

		return
	}

	saltPassword, err := utils.GenerateSalt(&data.Password)
	if err != nil {
		slog.Error(err.Error())

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Create User failed",
		})

		return
	}

	data.Password = saltPassword
	if err := service.CreateUser(data.Account, data.Password, data.Role); err != nil {
		slog.Error(err.Error())

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Create User failed",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User create success",
	})
}

func UserLogin(c *gin.Context) {
	var data struct {
		Account  string `json:"account" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.BindJSON(&data); err != nil {
		slog.Error(err.Error())

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Parameter",
		})

		return
	}

	user, err := service.QueryUser(data.Account)
	if err != nil {
		slog.Error(err.Error())

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Incorrect Username or Password",
		})

		return
	}

	if !utils.CompareSalt([]byte(user.Password), &data.Password) {
		slog.Error("password is incorrect")

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Incorrect Username or Password",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User Login Success",
	})
}
