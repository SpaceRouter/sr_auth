package sr_auth

import "testing"

var auth Auth = Auth{Key: "JESUISHYPERFANDESEXE", AuthServerAddress: "http://localhost:8085"}

func TestCreateAuth(t *testing.T) {
	auth := CreateAuth("JESUISHYPERFANDESEXE", "http://localhost:8085", nil)
	err := auth.PingAuthServer()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetUserFromToken(t *testing.T) {
	user, err := auth.GetUserFromToken(token)
	if err != nil {
		t.Fatal(err)
	}

	if user.GetUsername() != "root" {
		t.Fatalf("input != output")
	}
}
