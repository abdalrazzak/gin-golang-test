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
  
	router.GET("/products", GetProducts)

	router.Run() // listen and serve on 0.0.0.0:8080
}
 
func GetProducts(c *gin.Context) {

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	db.AutoMigrate(&Product{})

	var products []Product

	if err := db.Find(&products).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, products)
		log.Println(products)
	}

}