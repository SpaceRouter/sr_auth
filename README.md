# sr_auth

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

	auth := sr_auth.CreateAuth(key)
	user, err := auth.GetUserFromToken(tokenString)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(user)
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

	auth := sr_auth.CreateAuth(key)
	
	router.Use(auth.SrAuthMiddleware())

	router.GET("/info", func(c *gin.Context) {
		info, exist := c.Get("user")
		if !exist {
			c.AbortWithStatus(500)
			return
		}
		c.JSON(200, info)
	})
}
```