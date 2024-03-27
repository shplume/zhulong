package controller

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shplume/zhulong/service"
)

func CreateUser(c *gin.Context) {
	var data struct {
		Account  string `json:"account"`
		Password string `json:"password"`
		Role     int    `json:"role"`
	}

	if err := c.BindJSON(&data); err != nil {
		slog.Error(err.Error())

		c.JSON(http.StatusBadRequest, &gin.H{
			"message": "Invalid Parameter",
		})

		return
	}

	if err := service.CreateUser(data.Account, data.Password, data.Role); err != nil {
		slog.Error(err.Error())

		c.JSON(http.StatusInternalServerError, &gin.H{
			"message": err,
		})

		return
	}

	c.JSON(http.StatusOK, &gin.H{
		"message": "User create success",
	})
}
