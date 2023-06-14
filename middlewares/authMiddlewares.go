package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	helpers "github.com/goofynugtz/gin-jwt-auth-boilerplate/helpers"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization")
		splitToken := strings.Split(authorization, "Bearer ")
		accessToken := splitToken[1]
		if accessToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No Authorization header provided"})
			c.Abort()
			return
		}

		claims, err := helpers.ValidateToken(accessToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}
		c.Set("email", claims.Email)
		c.Set("first_name", claims.FirstName)
		c.Set("last_name", claims.LastName)
		c.Set("uid", claims.Uid)
		c.Set("user_type", claims.UserType)
		c.Next()
	}
}
