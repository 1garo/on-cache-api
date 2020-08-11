package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"log"
)

type LOGIN struct {
	USER     string `json:"user" binding:"required"`
	PASSWORD string `json:"password" binding:"required"`
	ID 			 int 		`json:"id"`
}

func main() {
	login := []*LOGIN{}
	var i int = 0
	var ids []int
	log.Printf("%d", ids)
	r := gin.Default()
	r.GET("/api/:id", func(c *gin.Context) {
		var loginTemp *LOGIN
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Fatalf("Cannot convert the id passed! \n%s", err)
			
		}
		// TODO: Make 400 and 500 handlers
		if err := c.ShouldBindJSON(&loginTemp); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		log.Printf("%s", c.Request.Write)
		c.JSON(http.StatusOK, gin.H{
			"message": login[id].ID,
			})
		})
	r.POST("/api", func(c *gin.Context) {
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
