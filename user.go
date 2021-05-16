package sr_auth

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
)

type Role string

type User struct {
	Login     string
	FirstName string
	LastName  string
	Email     string
}

type userRolesResponse struct {
	Message string
	Ok      bool
	Roles   []Role
}

type UserInfoResponse struct {
	Message string
	Ok      bool
	User    User
}

func (auth *Auth) GetRoles(token string) ([]Role, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: auth.CheckTLS},
	}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("GET", auth.AuthServerAddress+"/v1/roles", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var buffer bytes.Buffer
	_, err = buffer.ReadFrom(resp.Body)
	if err != nil {
		return nil, err
	}

	var rolesResponse userRolesResponse
	err = json.Unmarshal(buffer.Bytes(), &rolesResponse)
	if err != nil {
		return nil, err
	}

	if !rolesResponse.Ok {
		return nil, fmt.Errorf("auth server error : %s", rolesResponse.Message)
	}

	return rolesResponse.Roles, nil
}

func (auth *Auth) GetUserInfo(token string) (*User, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: auth.CheckTLS},
	}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("GET", auth.AuthServerAddress+"/v1/info", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var buffer bytes.Buffer
	_, err = buffer.ReadFrom(resp.Body)
	if err != nil {
		return nil, err
	}

	var rolesResponse UserInfoResponse
	err = json.Unmarshal(buffer.Bytes(), &rolesResponse)
	if err != nil {
		return nil, err
	}

	if !rolesResponse.Ok {
		return nil, fmt.Errorf("auth server error : %s", rolesResponse.Message)
	}

	return &rolesResponse.User, nil
}

func HasRole(roles []Role, role Role) bool {
	for _, uRole := range roles {
		if role == uRole {
			return true
		}
	}
	return false
}
