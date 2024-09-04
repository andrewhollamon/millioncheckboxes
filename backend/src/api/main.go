package main

import (
	"fmt"
	"github.com/andrewhollamon/millioncheckboxes/api/memoryStore"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	err := r.SetTrustedProxies(nil)
	if err != nil {
		fmt.Printf("Error trying to disabled trusted proxies: %v", err)
		os.Exit(1)
	}

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.POST("/api/v1/checkbox/:checkboxNbr/check", checkboxCheck)
	r.POST("/api/v1/checkbox/:checkboxNbr/uncheck", checkboxUncheck)

	return r
}

func main() {

	// first setup the memory store
	memoryStore.Init()

	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	err := r.Run(":8080")
	if err == nil {
		fmt.Printf("Router ended with error %v", err)
		return
	}
}
