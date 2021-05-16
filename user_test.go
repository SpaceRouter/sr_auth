package sr_auth

import "testing"

var token string = "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6InRlc3QiLCJleHAiOjE2MjEyODk4NjcsImlzcyI6InNwYWNlcm91dGVyIn0.H3yD-zOf7_64sDFNwwiq9cDCTYLr315XNZfy5K2UOmGHiqTN5TSRkoBRiIozFQbSGCrVYKodEurHaryHnhTteQ"

func TestGetRoles(t *testing.T) {
	roles, err := auth.GetRoles(token)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(roles)
}

func BenchmarkGetRoles(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := auth.GetRoles(token)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func TestAuth_GetUserInfo(t *testing.T) {
	infos, err := auth.GetUserInfo(token)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(infos)
}

func BenchmarkAuth_GetUserInfo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := auth.GetUserInfo(token)
		if err != nil {
			b.Fatal(err)
		}
	}
}
