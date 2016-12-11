package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"./database"
)

func main() {
	// Open the my.db data file in your current directory.
	// It will be created if it already doesn't exist.
	boltDB, _ := database.OpenDB()
	database.InitDB(boltDB)
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Connection Established")
	})

	router.GET("/test/:arg", func(c *gin.Context) {
		param := c.Param("arg")
		c.String(http.StatusOK, "Hello %s", param)
	})

	router.POST("/save/:key/:value", func(c *gin.Context) {
		key := c.Param("key")
		value := c.Param("value")
		var response *database.Response
		response = database.WriteSomethingToDB(boltDB, key, value)
		if response.Value == "error" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "error",
			})
		} else {
			c.JSON(http.StatusOK, response)
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

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
