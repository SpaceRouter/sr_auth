package sr_auth

import (
	"encoding/json"
	"testing"
)

var token string = "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6InJvb3QiLCJleHAiOjE2MjMyNDA1OTEsImlhdCI6MTYyMzE1NDE5MSwiaXNzIjoic3BhY2Vyb3V0ZXIifQ.3OvbLtQbJanWvyZ3fnGAZPbp3zbJH7Moycg4PSSliD8rlHP72e_KCF8Vq6mhuTkDATTiu9RqV5KQ9U_YEu0jew"

func TestGetRoles(t *testing.T) {
	user, err := auth.GetUserFromToken(token)
	if err != nil {
		t.Fatal(err)
	}
	roles, err := user.GetRoles()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(roles)
}

func BenchmarkGetRoles(b *testing.B) {
	user, err := auth.GetUserFromToken(token)
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		_, err := user.GetRoles()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func TestAuth_GetUserInfo(t *testing.T) {
	user, err := auth.GetUserFromToken(token)
	if err != nil {
		t.Fatal(err)
	}
	infos, err := user.GetUserInfo()
	if err != nil {
		t.Fatal(err)
	}
	formatted, err := json.Marshal(infos)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(formatted))
}

func BenchmarkAuth_GetUserInfo(b *testing.B) {
	user, err := auth.GetUserFromToken(token)
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		_, err := user.GetUserInfo()
		if err != nil {
			b.Fatal(err)
		}
	}
}
