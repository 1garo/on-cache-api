package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"on-cache-api/models"
	"strconv"
	"sync/atomic"
)

// TODO: create a ok in functions that don't return an err
var _id = new(int32)

func GetAllUsers(c *gin.Context) {
	var login []*models.LOGIN
	login = models.GetUsers()
	for _, value := range login {
		log.Printf("GetAllUsers(): %v : %d", value, *_id)
	}

	if login == nil {
		c.JSON(http.StatusOK, gin.H{
			"err": "don't exist any user in the memory!",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": login,
		})
	}
}

func GetUserBySHA(c *gin.Context) {
	// implement header auth check the sha
	// auth := c.Request.Header["Authorization"][0]
	sha := c.Param("sha")
	login := models.GetLogin()
	code, id, ok := models.CheckUserWithSHA(sha, login)
	if !ok {
		c.JSON(code, gin.H{
			"message": "user doesn't exist in memory!",
		})
	} else {
		c.JSON(code, gin.H{
			"id":    login[id].ID,
			"user":  login[id].USER,
			"email": login[id].EMAIL,
		})
	}
}

func GetDataByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatalf("Cannot convert the id passed! \n%s", err)

	}
	login := models.GetLogin()
	log.Printf("%d id da request\n", id)
	code, ok := models.CheckId(id, login)
	if !ok {
		c.JSON(code, gin.H{
			"message": "id does't exist in memory!",
		})
	} else {
		c.JSON(code, gin.H{
			"user":  login[id].USER,
			"email": login[id].EMAIL,
		})
	}
}

func SetUser(c *gin.Context) {
	// TODO: call the create sha function
	var loginTemp *models.LOGIN
	var login []*models.LOGIN
	err := c.BindJSON(&loginTemp)
	if err != nil {
		log.Fatalf("%s", err)
	}
	login = models.SetLogin(loginTemp, _id)
	log.Printf("SetData(): %v : %d", login[*_id], *_id)
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("User %s was added with success!", login[*_id].USER),
		"hash":    login[*_id].HASH,
	})
	atomic.AddInt32(_id, 1)
	log.Printf("SetData(): %d", *_id)
}
