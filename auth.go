package sr_auth

import (
	"github.com/dgrijalva/jwt-go"
)

type Auth struct {
	Key               string
	AuthServerAddress string
}

type customClaims struct {
	Username string
	jwt.StandardClaims
}

func CreateAuth(key string, authServerAddress string) *Auth {
	return &Auth{Key: key, AuthServerAddress: authServerAddress}
}

func (auth *Auth) GetUsernameFromToken(token string) (string, error) {
	tokenParsed, err := jwt.ParseWithClaims(token, &customClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(auth.Key), nil
	})

	if err != nil {
		return "", err
	}

	return tokenParsed.Claims.(*customClaims).Username, nil
}
