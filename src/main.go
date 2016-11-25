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
	database.OpenDB()
	router := gin.Default()

	router.GET("/:arg", func(c *gin.Context) {
		param := c.Param("arg")
		c.String(http.StatusOK, "Hello %s", param)
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
