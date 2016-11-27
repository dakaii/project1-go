package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"./database"
)

func main() {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	//boltDB = database.BoltDB
	boltDB, _ := database.OpenDB()
	database.InitDB(boltDB)
	//database.WriteSomethingToDB(boltDB)
	//database.RetrieveSomethingFromDB(boltDB)
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Connection Established")
	})

	router.GET("/test/:arg", func(c *gin.Context) {
		param := c.Param("arg")
		c.String(http.StatusOK, "Hello %s", param)
	})

	router.POST("/save/:arg", func(c *gin.Context) {
		key := c.Param("arg")
		value := database.WriteSomethingToDB(boltDB, key)
		if value != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"key":   key,
				"error": value,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"key": key,
			})
		}
	})

	router.GET("/retrieve/:arg", func(c *gin.Context) {
		key := c.Param("arg")
		value := database.RetrieveSomethingFromDB(boltDB, key)
		c.JSON(http.StatusOK, gin.H{
			"key":   key,
			"value": value,
		})
	})

	// Simple group: v1
	//cache := router.Group("/cache")
	//{
	//cache.POST("/login", database.##)
	//cache.GET("/get", database.@@)
	//}

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
