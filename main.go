package main

import (
	"chipmonk_assignment/routes"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	if startErr := startServer(); startErr != nil {
		log.Fatal(startErr)
	}
}

// startServer -  to start our gin server
func startServer() error {
	r := gin.Default()
	r.GET("/", checkStatus())
	routes.InitRoutes(r)
	s := &http.Server{
		Handler: r,
		Addr:    ":7000",
	}
	if startErr := s.ListenAndServe(); startErr != nil {
		return startErr
	}
	return nil
}

// checkStatus -  to check the server running status
func checkStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "Server Running Successfully")
	}
}
