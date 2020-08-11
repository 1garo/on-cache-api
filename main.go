package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LOGIN struct {
	USER     string `json:"user" binding:"required"`
	PASSWORD string `json:"password" binding:"required"`
}

func main() {
  
	login := []*login{}
	var i int = 0
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": login,
		})
	})
	r.POST("/foo", func(c *gin.Context) {
		var loginTemp *LOGIN
		c.BindJSON(&loginTemp)
		login = append(login, loginTemp)
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%s was added with success!", login[i].USER),
		})
		i += 1
	})
	r.Run()
}
