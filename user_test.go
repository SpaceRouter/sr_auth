package sr_auth

import (
	"encoding/json"
	"testing"
)

var token string = "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6InRlc3QiLCJleHAiOjE2MjEyODk4NjcsImlzcyI6InNwYWNlcm91dGVyIn0.H3yD-zOf7_64sDFNwwiq9cDCTYLr315XNZfy5K2UOmGHiqTN5TSRkoBRiIozFQbSGCrVYKodEurHaryHnhTteQ"

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
