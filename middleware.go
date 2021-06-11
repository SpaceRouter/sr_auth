package sr_auth

import (
	"fmt"
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
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "SrAuthMiddleware Error : " + err.Error(), "Ok": false})
			return
		}

		c.Set("user", u)
		c.Next()
	}
}

func ExtractUser(c *gin.Context) (*User, error) {
	userObject, exist := c.Get("user")
	if !exist {
		return nil, fmt.Errorf("user can't be found in session")
	}

	return userObject.(*User), nil
}

func (auth *Auth) SrAuthHttp(r *http.Request) (*User, error) {
	tokenString := r.Header.Get("Authorization")
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	tokenString = strings.Replace(tokenString, "bearer ", "", 1)

	return auth.GetUserFromToken(tokenString)
}
