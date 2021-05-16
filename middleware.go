package sr_auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (auth *Auth) SrAuthMiddlewareGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		tokenString = strings.Replace(tokenString, "bearer ", "", 1)

		u, err := auth.GetUserFromToken(tokenString)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
			return
		}

		c.Set("user", u)
		c.Next()
	}
}

func (auth *Auth) SrAuthHttp(r *http.Request) (User, error) {
	tokenString := r.Header.Get("Authorization")
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	tokenString = strings.Replace(tokenString, "bearer ", "", 1)

	return auth.GetUserFromToken(tokenString)
}
