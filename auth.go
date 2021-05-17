package sr_auth

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

func CreateAuth(key string, authServerAddress string, tlsConfig *tls.Config) *Auth {
	return &Auth{Key: key, AuthServerAddress: authServerAddress, TlsConfig: tlsConfig}
}

func (auth *Auth) GetUserFromToken(token string) (*User, error) {
	tokenParsed, err := jwt.ParseWithClaims(token, &customClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(auth.Key), nil
	})

	if err != nil {
		return nil, err
	}

	return &User{username: tokenParsed.Claims.(*customClaims).Username, token: token, auth: auth}, nil
}

func (auth *Auth) PingAuthServer() error {
	tr := &http.Transport{
		TLSClientConfig: auth.TlsConfig,
	}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("GET", auth.AuthServerAddress+"/health", nil)
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	var buffer bytes.Buffer
	_, err = buffer.ReadFrom(resp.Body)
	if err != nil {
		return err
	}
	ret := string(buffer.Bytes())

	if ret != "Ok" {
		return fmt.Errorf("replied \"%s\" instead of \"Ok\"", ret)
	}

	return nil
}
