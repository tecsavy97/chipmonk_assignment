package handlers

import (
	"chipmonk_assignment/models"
	"chipmonk_assignment/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// RegisterUser -  Handler to handle registration route
func RegisterUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := models.NewUser()
		if bindErr := c.Bind(&user); bindErr != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if regUser := services.RegisterUser(user); regUser != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": regUser.Error()})
			return
		}
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "User Registered Successfully!!"})
		return
	}
}

// LoginUser - Handler to handle login route
func LoginUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := models.NewUser()
		if bindErr := c.Bind(&user); bindErr != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		token, logUser := services.LoginUser(user)
		if logUser != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": logUser.Error()})
			return
		}
		c.Header("Authorization", token)
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "User LoggedIn. Please user the following token in request header", "token": token})
		return
	}
}

// FetchAllActiveUser -  Handler to handle fetch user route
func FetchAllActiveUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" && strings.TrimSpace(token) == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Please set the given token as `Authorization` in header"})
			return
		}
		data, fetchUserErr := services.FetchActiveUsers(token)
		if fetchUserErr != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": fetchUserErr.Error()})
			return
		}
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"userlist": data})
		return

	}
}
