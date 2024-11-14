package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

const URI = "mongodb://localhost:27017"

func main() {

	client := CreateClient(URI)
	dbList := getDbList(client)

	for _, dbName := range dbList {
		fmt.Println(dbName)
	}
	r := gin.Default()
	routes := r.Group("/api")
	{
		routes.GET("/", getHome)
		routes.GET("/user/:name", getUser)
	}

	database := r.Group("/api/base")
	{
		database.GET("/a", getAll)
	}

	r.Run(":5000")

}

func getAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "CVE 234-as",
	})
}

func getHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to the API!",
	})
}

func getUser(c *gin.Context) {

	name := c.Param("name")
	c.JSON(200, gin.H{
		"message": "Hello, " + name + "!",
	})
}
