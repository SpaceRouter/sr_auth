package sr_auth

import (
	"crypto/tls"
	"github.com/dgrijalva/jwt-go"
)

type Role string

type UserInfo struct {
	Login     string
	FirstName string
	LastName  string
	Email     string
}

type Auth struct {
	Key               string
	AuthServerAddress string
	TlsConfig         *tls.Config
}

type customClaims struct {
	Username string
	jwt.StandardClaims
}

type User struct {
	username string
	token    string
	auth     *Auth
}
