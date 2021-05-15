package sr_auth

import (
	"github.com/dgrijalva/jwt-go"
)

type Auth struct {
	Key string
}

func CreateAuth(key string) Auth {
	return Auth{key}
}

type customClaims struct {
	User User
	jwt.StandardClaims
}

func (auth *Auth) GetUserFromToken(token string) (User, error) {
	tokenParsed, err := jwt.ParseWithClaims(token, &customClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(auth.Key), nil
	})

	if err != nil {
		return User{}, err
	}

	return tokenParsed.Claims.(*customClaims).User, nil
}
