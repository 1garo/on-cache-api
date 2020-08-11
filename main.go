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


func checkId (id int, login []*LOGIN) (int) {
	var code int
	log.Printf("%v", login[[0])
	for i := 0; i < len(login); i++ {
		if tempId := login[i].ID; tempId == id {
			code = http.StatusOK
		}else {
			code = http.StatusBadRequest
		}	
	}
	return code
}
// POST:  curl -H "Content-Type: application/json" -X POST -d '{"user":"thoma","password":"bobao"}' http://localhost:8080/api
// GET:  curl -H "Content-Type: application/json" -X GET http://localhost:8080/api/<id>
func main() {
	login := []*LOGIN{}
	var i int = 0
	var ids []int
	log.Printf("%d", ids)
	r := gin.Default()
	r.GET("/api/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Fatalf("Cannot convert the id passed! \n%s", err)
			
		}
		code := checkId(id, login)
		if code != 200 {
			c.JSON(code, gin.H{
				"message": "id does't exist in memory",
			})
		}else{
			c.JSON(code, gin.H{
				"user": login[id].USER,
				"password": login[id].PASSWORD,
				})
			}
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
