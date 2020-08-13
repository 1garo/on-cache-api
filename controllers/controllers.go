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

var _id = new(int32)

func checkUser(user string, login []*models.LOGIN) (int, int) {
	var code, id int
	log.Println(len(login))
	for i := 0; i < len(login); i++ {
		log.Printf("checkId(): %v %d", login[i], i)
		log.Printf("login[i].ID: %d", login[i].ID)
		if tempUser := login[i].USER; tempUser == user{
			code = http.StatusOK
			id = login[i].ID
			break
		} else {
			code = http.StatusBadRequest
			id = -1
		}
	}
	return code, id
}

func checkId(id int, login []*models.LOGIN) int {
	var code int
	log.Println(len(login))
	for i := 0; i < len(login); i++ {
		log.Printf("checkId(): %v %d", login[i], i)
		log.Printf("login[i].ID: %d", login[i].ID)
		if tempId := login[i].ID; tempId == id {
			code = http.StatusOK
			break
		} else {
			code = http.StatusBadRequest
		}
	}
	return code
}

func GetUserByID(c *gin.Context) {
	user := c.Param("user")
	login := models.GetLogin()
	code, id := checkUser(user, login)
	log.Printf("%d id da request\n", id)
	if id == -1 {
		c.JSON(code, gin.H{
			"message": "user doesn't exist in memory!",
		})
	} else {
		c.JSON(code, gin.H{
			"id": login[id].ID,
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
	code := checkId(id, login)
	if code != 200 {
		c.JSON(code, gin.H{
			"message": "id does't exist in memory!",
		})
	} else {
		c.JSON(code, gin.H{
			"user":     login[id].USER,
			"password": login[id].PASSWORD,
		})
	}
}

func SetData(c *gin.Context) {
	var loginTemp *models.LOGIN
	err := c.BindJSON(&loginTemp)
	if err != nil {
		log.Fatalf("%s", err)
	}
	var login []*models.LOGIN
	login = models.SetLogin(loginTemp, _id)
	log.Printf("SetData(): %v : %d", login[*_id], *_id)
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s was added with success!", login[*_id].USER),
	})
	atomic.AddInt32(_id, 1)
	log.Printf("SetData(): %d", *_id)

}
