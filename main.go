package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"on-cache-api/controllers"
)

// POST:  curl -H "Content-Type: application/json" -X POST -d '{"user":"thoma","password":"bobao"}' http://localhost:8080/api
// GET:  curl -H "Content-Type: application/json" -X GET http://localhost:8080/api/<id>
func main() {
	r := gin.Default()
	_ = r.Group("req")
	{
		r.GET("/user/id/:id", controllers.GetDataByID)
		r.GET("/user/sha/:sha", controllers.GetUserBySHA)
		r.GET("/users", controllers.GetAllUsers)
	}
	_ = r.Group("resp")
	{
		r.POST("/data", controllers.SetUser)
	}

	err := r.Run()
	if err != nil {
		log.Fatalf("%s", err)
	}
}
