package sr_auth

import (
	"github.com/dgrijalva/jwt-go"
)

type Auth struct {
	Key               string
	AuthServerAddress string
	CheckTLS          bool
}

type customClaims struct {
	Username string
	jwt.StandardClaims
}

func CreateAuth(key string, authServerAddress string, checkTLS bool) *Auth {
	return &Auth{Key: key, AuthServerAddress: authServerAddress, CheckTLS: true}
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
