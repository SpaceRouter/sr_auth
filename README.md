# sr_auth

[![](https://goreportcard.com//badge/github.com/SpaceRouter/sr_auth)](https://goreportcard.com/report/github.com/SpaceRouter/sr_auth)

## Example

## Vanilla usage

```go
package main

import (
	"github.com/spacerouter/sr_auth"
	"log"
)

func main() {
	key := "SECRETKEY"

	auth := sr_auth.CreateAuth(key, "http://localhost:8080", nil)
	user, err := auth.GetUserFromToken(tokenString)
	if err != nil {
		log.Fatal(err)
	}

	roles, err := user.GetRoles()
	if err != nil {
		return
	}

	log.Println(user)
	log.Println(roles)
}
```

- **key** is the token's secret key
- **user** contains user's information

## Gin usage

```go
package main

import (
	"github.com/spacerouter/sr_auth"
	"github.com/gin-gonic/gin"
)

func main() {
	key := "SECRETKEY"

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	auth := sr_auth.CreateAuth(key, "http://localhost:8080", nil)

	router.Use(auth.SrAuthMiddlewareGin())

	router.GET("/info", func(c *gin.Context) {
		userObject, exist := c.Get("user")
		if !exist {
			c.AbortWithStatus(500)
			return
		}

		user := userObject.(*sr_auth.User)

		roles, err := user.GetRoles()
		if err != nil {
			return
		}
		c.JSON(200, gin.H{"user": user, "roles": roles})
	})
}
```