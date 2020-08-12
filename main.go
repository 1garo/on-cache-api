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
	r.GET("/api/:id", controllers.GetDataByID)
	r.POST("/api", controllers.SetData)
	err := r.Run()
	if err != nil {
		log.Fatalf("%s", err)
	}
}
