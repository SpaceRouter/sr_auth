package sr_auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (u *User) GetUsername() string {
	return u.username
}

func (u *User) GetRoles() ([]Role, error) {
	tr := &http.Transport{
		TLSClientConfig: u.auth.TlsConfig,
	}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("GET", u.auth.AuthServerAddress+"/v1/roles", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+u.token)
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

func (u *User) GetUserInfo() (*UserInfo, error) {
	tr := &http.Transport{
		TLSClientConfig: u.auth.TlsConfig,
	}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("GET", u.auth.AuthServerAddress+"/v1/info", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+u.token)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var buffer bytes.Buffer
	_, err = buffer.ReadFrom(resp.Body)
	if err != nil {
		return nil, err
	}

	var rolesResponse userInfoResponse
	err = json.Unmarshal(buffer.Bytes(), &rolesResponse)
	if err != nil {
		return nil, err
	}

	if !rolesResponse.Ok {
		return nil, fmt.Errorf("auth server error : %s", rolesResponse.Message)
	}

	return &rolesResponse.UserInfo, nil
}

func HasRole(roles []Role, role Role) bool {
	for _, uRole := range roles {
		if role == uRole {
			return true
		}
	}
	return false
}
