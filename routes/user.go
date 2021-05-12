package routes

import (
	"chipmonk_assignment/handlers"

	"github.com/gin-gonic/gin"
)

// InitRoutes - to initialises all the routes for the APIs to be consumed
func InitRoutes(r *gin.Engine) {
	r.POST("/register", handlers.RegisterUser())
	r.POST("/login", handlers.LoginUser())
	r.GET("/active-users", handlers.FetchAllActiveUser())
}
