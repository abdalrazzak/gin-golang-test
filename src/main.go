package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
) 

func main() {

	router := gin.Default()


	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"greet": "hello, world!",
		})
	}) 

	router.Run() // listen and serve on 0.0.0.0:8080
}
  